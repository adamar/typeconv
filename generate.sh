#!/bin/bash



list="Int Int8 Int16 Int32 Int64 Uint Uint8 Uint16 Uint32 Uint64 Float32 Float64 Complex64 Complex128 String"
for output in $list; do 
    output_lower=$(echo $output | awk '{print tolower($0)}')
    filename="to_${output_lower}.go"
    >"${filename}"
    test_filename="to_${output_lower}_test.go"
    >"${test_filename}"

cat << EOF > "${filename}"

package typeconv

EOF


cat << EOF > "${test_filename}"

package typeconv

import (
    "testing"
)

EOF


    for input in $list; do
      if [[ "$input" != "$output" ]]; then

cat << EOF >> "${filename}"
// ${input}To${output} as the name implies takes a ${input} and converts it into a ${output}
// func ${input}To${output}(input $input) (${output}, error) {
//
//   NOT IMPLEMENTED YET
//
//}


EOF

cat << EOF >> "${test_filename}"

func Test${input}To${output}(t *testing.T) {

    input := ${input}(1)
    expected := ${output}(1)

    result, err := ${input}To${output}(input)

    if err != nil {
        t.Errorf("Error %v", err)
    }

    if result != expected {
        t.Errorf("Result was incorrect, got: %s, want: %s.", result, "Foo")
    }

}

EOF

      fi
    done
done

