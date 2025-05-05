package yamlparser

import (
	"gopkg.in/yaml.v3"
)

// ParseYAML парсит YAML данные в map[string]interface{}
func ParseYAML(data []byte) (map[string]interface{}, error) {
	var root yaml.Node
	err := yaml.Unmarshal(data, &root)
	if err != nil {
		return nil, err
	}

	result := make(map[string]interface{})
	parseToMap(&root, result)

	return result, nil
}

func parseToMap(node *yaml.Node, m map[string]interface{}) {
	switch node.Kind {
	case yaml.DocumentNode:
		for _, n := range node.Content {
			parseToMap(n, m)
		}
	case yaml.MappingNode:
		for i := 0; i < len(node.Content); i += 2 {
			key := node.Content[i]
			value := node.Content[i+1]

			switch value.Kind {
			case yaml.ScalarNode:
				m[key.Value] = value.Value
			case yaml.MappingNode:
				nested := make(map[string]interface{})
				parseToMap(value, nested)
				m[key.Value] = nested
			case yaml.SequenceNode:
				list := parseSequence(value)
				m[key.Value] = list
			}
		}
	}
}

func parseSequence(node *yaml.Node) []interface{} {
	var list []interface{}
	for _, n := range node.Content {
		switch n.Kind {
		case yaml.ScalarNode:
			list = append(list, n.Value)
		case yaml.MappingNode:
			m := make(map[string]interface{})
			parseToMap(n, m)
			list = append(list, m)
		case yaml.SequenceNode:
			list = append(list, parseSequence(n))
		}
	}
	return list
}
