add:
message MyMessage {
    int32 id = 1;
}

message MyMessage {
    int32 id = 1;
    string first_name = 2;
}

update:
message MyMessage {
    int32 id = 1;
    string first_name = 2;
}

message MyMessage {
    int32 id = 1;
    string new_first_name = 2;
}


remove:
message MyMessage {
    int32 id = 1;
    string first_name = 2;
}

message MyMessage {
    reserved 2;
    reserved "first_name";
    int32 id = 1;
}

!!!reserved prevent form re-using
!!!do not ever remove any reserved tags or fields
    tags:
    reserved 2, 15, 9 to 11;
    fields:
    reserved "name1", "name2";



---------------------------------
https://developers.google.com/protocol-buffers/docs/proto3
https://www.udemy.com/course/protocol-buffers/learn/lecture/9924762