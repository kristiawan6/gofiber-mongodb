package utils

import (
    "bytes"
    "encoding/binary"
    "gofiber-mongodb/src/models"
)

func EncodeBSON(data []models.Student) ([]byte, error) {
    // Create a bytes.Buffer to hold the BSON data
    buf := bytes.Buffer{}

    // Iterate over each student and encode them to BSON
    for _, student := range data {
        // Write name length
        nameLen := int32(len(student.Name))
        err := binary.Write(&buf, binary.LittleEndian, nameLen)
        if err != nil {
            return nil, err
        }
        // Write name
        _, err = buf.WriteString(student.Name)
        if err != nil {
            return nil, err
        }
        // Write grade
        err = binary.Write(&buf, binary.LittleEndian, int32(student.Grade))
        if err != nil {
            return nil, err
        }
    }

    return buf.Bytes(), nil
}
