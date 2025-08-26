//go:build ignore

// Package main is a generator for data-to-domain converters.
package main

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Field struct {
	DomainName string
	Converter  string
}

func main() {
	const tpl = `
package mapper

import (
	"github.com/go-kratos/kratos-layout/internal/bondProductApi/biz/domain"
	model "github.com/go-kratos/kratos-layout/internal/bondProductApi/data/grom/models"
	"github.com/shopspring/decimal"
)


// decimalToString converts decimal.NullDecimal to string.
func decimalToString(d decimal.NullDecimal) string {
	if d.Valid {
		return d.Decimal.String()
	}
	return ""
}

// ToBondProductItem converts model.TProduct to domain.BondProductItem.
func ToBondProductItem(do *model.TProduct) *domain.BondProductItem {
	if do == nil {
		return nil
	}
	return &domain.BondProductItem{
		{{- range .Fields}}
		{{.DomainName}}: {{.Converter}},{{- end}}
	}
}

`

	fset := token.NewFileSet()
	filePath := "../models/t_product.go"
	abs, _ := filepath.Abs(filePath)
	fmt.Println("Parsing:", abs)

	node, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		log.Fatal("ParseFile error:", err)
	}

	var fields []Field
	var structName = "TProduct"

	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok || genDecl.Tok != token.TYPE {
			continue
		}
		for _, spec := range genDecl.Specs {
			tspec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}
			if tspec.Name.Name != structName {
				continue
			}

			// 断言是否为结构体
			structType, ok := tspec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			// 遍历结构体字段
			for _, field := range structType.Fields.List {
				if len(field.Names) == 0 {
					continue
				}
				name := field.Names[0].Name

				// 解析 tag
				var jsonTag string
				if field.Tag != nil {
					jsonTag = parseJSONTag(field.Tag.Value)
				}
				if jsonTag == "-" {
					continue
				}

				// 默认使用字段名
				domainName := name
				converter := fmt.Sprintf("do.%s", name)

				// 特殊类型处理
				if isDecimalNullDecimal(field.Type) {
					converter = fmt.Sprintf("decimalToString(do.%s)", name)
				}

				fields = append(fields, Field{
					DomainName: domainName,
					Converter:  converter,
				})
			}
			break
		}
	}

	if len(fields) == 0 {
		log.Fatal("No fields found in TProduct")
	}

	// 创建输出目录
	outputDir := "../mapper"
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		log.Fatal("MkdirAll error:", err)
	}

	// 创建输出文件
	outputFile := filepath.Join(outputDir, "product_gen.go")
	file, err := os.Create(outputFile)
	if err != nil {
		log.Fatal("Create file error:", err)
	}
	defer file.Close()

	// 执行模板
	t := template.Must(template.New("converter").Parse(tpl))
	data := map[string]interface{}{
		"Fields": fields,
		"Module": getModulePath(), // 自动获取模块名
	}

	if err := t.Execute(file, data); err != nil {
		log.Fatal("Template execute error:", err)
	}

	fmt.Printf("✅ Generated: %s\n", outputFile)
}

// parseJSONTag 解析 json tag
func parseJSONTag(tag string) string {
	tag = strings.Trim(tag, "`")
	for _, part := range strings.Split(tag, " ") {
		if strings.HasPrefix(part, "json:") {
			content := strings.TrimPrefix(part, "json:")
			content = strings.Trim(content, `"`)
			if idx := strings.Index(content, ","); idx != -1 {
				content = content[:idx]
			}
			return content
		}
	}
	return ""
}

// isDecimalNullDecimal 判断是否是 decimal.NullDecimal
func isDecimalNullDecimal(expr ast.Expr) bool {
	if sel, ok := expr.(*ast.SelectorExpr); ok {
		if x, ok := sel.X.(*ast.Ident); ok {
			return x.Name == "decimal" && sel.Sel.Name == "NullDecimal"
		}
	}
	return false
}

// getModulePath 尝试读取 go.mod 获取模块名
func getModulePath() string {
	data, err := os.ReadFile("../../../../go.mod")
	if err != nil {
		return "your-module"
	}
	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "module ") {
			return strings.TrimSpace(strings.TrimPrefix(line, "module"))
		}
	}
	return "your-module"
}
