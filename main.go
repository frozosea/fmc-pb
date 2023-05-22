package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"

	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/desc/builder"
	"github.com/jhump/protoreflect/desc/protoparse"
	"github.com/jhump/protoreflect/desc/protoprint"
)

func ConcatFiles(importPaths []string, fileNames ...string) (io.Reader, error) {
	files, err := parse(importPaths, fileNames...)
	if err != nil {
		return nil, err
	}
	rd, err := concat(files...)
	if err != nil {
		return nil, err
	}
	b, err := io.ReadAll(rd)
	if err != nil {
		return nil, err
	}
	f := string(b)
	// remove duplicated imports
	for _, fn := range fileNames {
		f = strings.ReplaceAll(f, fmt.Sprintf(`import "%s";
`, fn), "")
	}
	return strings.NewReader(f), nil
}

func parse(importPaths []string, fileNames ...string) ([]*desc.FileDescriptor, error) {
	p := protoparse.Parser{
		ImportPaths:           importPaths,
		InferImportPaths:      true,
		IncludeSourceCodeInfo: true,
	}
	files, err := p.ParseFiles(fileNames...)
	if err != nil {
		return nil, fmt.Errorf("failed to parse proto files: %w", err)
	}
	return files, nil
}

func concat(files ...*desc.FileDescriptor) (io.Reader, error) {
	if len(files) == 0 {
		return strings.NewReader(""), nil
	}

	b := builder.NewFile("result.proto")

	firstFile := files[0]
	isProto3 := firstFile.IsProto3()
	pkg := firstFile.GetPackage()

	b.SetProto3(isProto3)
	b.SetPackageName(pkg)

	for _, file := range files {
		if file.GetPackage() != pkg {
			return nil, fmt.Errorf("package name want: %s, got: %s", pkg, file.GetPackage())
		}
		if file.IsProto3() != isProto3 {
			return nil, fmt.Errorf("isProto3 want: %b, got: %b", isProto3, file.IsProto3())
		}

		for _, m := range file.GetMessageTypes() {
			mb, err := fromMessage(m)
			if err != nil {
				return nil, err
			}
			b.AddMessage(mb)
		}
		for _, e := range file.GetEnumTypes() {
			eb, err := builder.FromEnum(e)
			if err != nil {
				return nil, err
			}
			b.AddEnum(eb)
		}
		for _, ex := range file.GetExtensions() {
			exb, err := builder.FromField(ex)
			if err != nil {
				return nil, err
			}
			b.AddExtension(exb)
		}
		for _, s := range file.GetServices() {
			sb, err := builder.FromService(s)
			if err != nil {
				return nil, err
			}
			b.AddService(sb)
		}
	}

	f, err := b.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build concatenated proto: %w", err)
	}
	printer := protoprint.Printer{
		SortElements: true,
	}
	var buf bytes.Buffer
	err = printer.PrintProtoFile(f, &buf)
	if err != nil {
		return nil, fmt.Errorf("failed to print proto: %w", err)
	}
	return &buf, nil
}

func fromMessage(m *desc.MessageDescriptor) (*builder.MessageBuilder, error) {
	mProto := m.AsDescriptorProto()
	mb := builder.NewMessage(m.GetName())
	c := builder.Comments{
		LeadingDetachedComments: m.GetSourceInfo().GetLeadingDetachedComments(),
		LeadingComment:          m.GetSourceInfo().GetLeadingComments(),
		TrailingComment:         m.GetSourceInfo().GetTrailingComments(),
	}
	mb.SetComments(c)
	mb.SetExtensionRanges(mProto.GetExtensionRange())
	mb.SetOptions(mProto.GetOptions())
	mb.SetReservedNames(mProto.GetReservedName())
	for _, f := range m.GetFields() {
		fld, err := builder.FromField(f)
		if err != nil {
			return nil, err
		}
		mb.AddField(fld)
	}
	for _, childMsg := range m.GetNestedMessageTypes() {
		msg, err := builder.FromMessage(childMsg)
		if err != nil {
			return nil, err
		}
		mb.AddNestedMessage(msg)
	}
	for _, childEnum := range m.GetNestedEnumTypes() {
		enum, err := builder.FromEnum(childEnum)
		if err != nil {
			return nil, err
		}
		mb.AddNestedEnum(enum)
	}
	return mb, nil
}

func contains(s []string, v string) bool {
	for _, i := range s {
		if i == v && v != "}" && v != "{" {
			return true
		}
	}
	return false
}

func main() {
	uF, err := ioutil.ReadFile("user/user.proto")
	if err != nil {
		log.Fatal(err)
	}
	u := strings.Split(string(uF), "package user;")[1]

	tF, err := ioutil.ReadFile("tracking/tracking.proto")
	if err != nil {
		log.Fatal(err)
	}
	t := strings.Replace(string(tF), "package tracking;", "package fmc;", 1)

	sF, err := ioutil.ReadFile("schedule-tracking/schedule_tracking.proto")
	//fmt.Println(strings.Split(string(sF), "package schedule_tacking;\n")[])
	s := strings.Split(string(sF), "package schedule_tacking;\n")[1]

	var arr []string

	full := "syntax = \"proto3\";\n\npackage fmc;" + "\n\n" + u + "\n\n" + t + "\n\n" + s
	for _, v := range strings.Split(full, "\n") {
		if !contains(arr, v) {
			arr = append(arr, v)
		}
	}
	if err := ioutil.WriteFile("fmc.proto", []byte(strings.Join(arr, "\n")), 0777); err != nil {
		log.Fatal(err)
	}
	//protoc --proto_path=${PWD} --go_out=${PWD}/tracking --go-grpc_out=${PWD}/tracking  --grpc-gateway_out=${PWD}/tracking  --grpc-gateway_opt allow_delete_body=true --grpc-gateway_opt generate_unbound_methods=true --openapiv2_out=allow_delete_body=true:${PWD}/tracking -I ${PWD}/tracking tracking.proto

	//const protoc = "protoc"
	//cmd := exec.Command(protoc, []string{"--proto_path=fmc.proto", "--go_out=./", "--grpc-gateway_out=./", "--grpc-gateway_opt", "allow_delete_body=true", "--grpc-gateway_opt", "generate_unbound_methods=true", "--openapiv2_out=allow_delete_body=true:./", "-I", "./fmc.proto"}...)
	//stdout, err := cmd.Output()
	//if err != nil {
	//	log.Fatal(err.Error())
	//	return
	//}
	//fmt.Println(string(stdout))
}
