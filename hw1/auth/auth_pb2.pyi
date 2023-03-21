from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Optional as _Optional

DESCRIPTOR: _descriptor.FileDescriptor

class AuthByNameRequest(_message.Message):
    __slots__ = ["name"]
    NAME_FIELD_NUMBER: _ClassVar[int]
    name: str
    def __init__(self, name: _Optional[str] = ...) -> None: ...

class AuthByNameResponse(_message.Message):
    __slots__ = ["authed"]
    AUTHED_FIELD_NUMBER: _ClassVar[int]
    authed: bool
    def __init__(self, authed: bool = ...) -> None: ...
