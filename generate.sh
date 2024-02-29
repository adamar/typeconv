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

//import (
//    "testing"
//)

EOF


    for input in $list; do
      if [[ "$input" != "$output" ]]; then

        input_lower=$(echo $input | awk '{print tolower($0)}')

cat << EOF >> "${filename}"
// ${input}To${output} as the name implies takes a ${input_lower} and converts it into a ${output_lower}
// func ${input}To${output}(input $input_lower) (${output_lower}, error) {

//   NOT IMPLEMENTED YET

//}


EOF

cat << EOF >> "${test_filename}"
//func Test${input}To${output}(t *testing.T) {

//    input := ${input_lower}(1)
//    expected := ${output_lower}(1)

//    result, err := ${input}To${output}(input)

//    if err != nil {
//        t.Errorf("Error %v", err)
//    }

//    if result != expected {
//        t.Errorf("Result was Incorrect, got: %s, wanted: %s.", result, expected)
//    }

//}


EOF

      fi
    done
done

