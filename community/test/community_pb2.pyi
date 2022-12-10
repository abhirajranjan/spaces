from google.protobuf.internal import containers as _containers
from google.protobuf.internal import enum_type_wrapper as _enum_type_wrapper
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor
ERROR: STATUS
INT32_DEFAULT: DEFAULTS
NOT_FOUND: STATUS
OK: STATUS
STRING_DEFAULT: DEFAULTS
WRONG_FORMAT: STATUS
one: NUMBER_NUM
three: NUMBER_NUM
two: NUMBER_NUM

class CommunityGetRequest(_message.Message):
    __slots__ = ["Name", "Tag", "id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    Name: str
    TAG_FIELD_NUMBER: _ClassVar[int]
    Tag: str
    id: int
    def __init__(self, id: _Optional[int] = ..., Name: _Optional[str] = ..., Tag: _Optional[str] = ...) -> None: ...

class CommunityMetaData(_message.Message):
    __slots__ = ["Banner", "Description", "Id", "Name", "Spaces", "Tag"]
    BANNER_FIELD_NUMBER: _ClassVar[int]
    Banner: str
    DESCRIPTION_FIELD_NUMBER: _ClassVar[int]
    Description: str
    ID_FIELD_NUMBER: _ClassVar[int]
    Id: int
    NAME_FIELD_NUMBER: _ClassVar[int]
    Name: str
    SPACES_FIELD_NUMBER: _ClassVar[int]
    Spaces: _containers.RepeatedCompositeFieldContainer[SpacesCompact]
    TAG_FIELD_NUMBER: _ClassVar[int]
    Tag: str
    def __init__(self, Id: _Optional[int] = ..., Name: _Optional[str] = ..., Tag: _Optional[str] = ..., Description: _Optional[str] = ..., Banner: _Optional[str] = ..., Spaces: _Optional[_Iterable[_Union[SpacesCompact, _Mapping]]] = ...) -> None: ...

class MStatus(_message.Message):
    __slots__ = ["Status"]
    STATUS_FIELD_NUMBER: _ClassVar[int]
    Status: STATUS
    def __init__(self, Status: _Optional[_Union[STATUS, str]] = ...) -> None: ...

class SpacesCompact(_message.Message):
    __slots__ = ["Id"]
    ID_FIELD_NUMBER: _ClassVar[int]
    Id: str
    def __init__(self, Id: _Optional[str] = ...) -> None: ...

class DEFAULTS(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []

class NUMBER_NUM(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []

class STATUS(int, metaclass=_enum_type_wrapper.EnumTypeWrapper):
    __slots__ = []
