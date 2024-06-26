from google.protobuf.internal import containers as _containers
from google.protobuf import descriptor as _descriptor
from google.protobuf import message as _message
from typing import ClassVar as _ClassVar, Iterable as _Iterable, Mapping as _Mapping, Optional as _Optional, Union as _Union

DESCRIPTOR: _descriptor.FileDescriptor

class Empty(_message.Message):
    __slots__ = ()
    def __init__(self) -> None: ...

class TimeSeriesItem(_message.Message):
    __slots__ = ("ts", "value")
    TS_FIELD_NUMBER: _ClassVar[int]
    VALUE_FIELD_NUMBER: _ClassVar[int]
    ts: int
    value: float
    def __init__(self, ts: _Optional[int] = ..., value: _Optional[float] = ...) -> None: ...

class MakePredictRequest(_message.Message):
    __slots__ = ("username", "name", "unit", "period", "predict_periods", "items")
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    UNIT_FIELD_NUMBER: _ClassVar[int]
    PERIOD_FIELD_NUMBER: _ClassVar[int]
    PREDICT_PERIODS_FIELD_NUMBER: _ClassVar[int]
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    username: str
    name: str
    unit: str
    period: int
    predict_periods: int
    items: _containers.RepeatedCompositeFieldContainer[TimeSeriesItem]
    def __init__(self, username: _Optional[str] = ..., name: _Optional[str] = ..., unit: _Optional[str] = ..., period: _Optional[int] = ..., predict_periods: _Optional[int] = ..., items: _Optional[_Iterable[_Union[TimeSeriesItem, _Mapping]]] = ...) -> None: ...

class GetPredictRequest(_message.Message):
    __slots__ = ("username", "name")
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    NAME_FIELD_NUMBER: _ClassVar[int]
    username: str
    name: str
    def __init__(self, username: _Optional[str] = ..., name: _Optional[str] = ...) -> None: ...

class GetPredictResponse(_message.Message):
    __slots__ = ("unit", "delimiter", "period", "items")
    UNIT_FIELD_NUMBER: _ClassVar[int]
    DELIMITER_FIELD_NUMBER: _ClassVar[int]
    PERIOD_FIELD_NUMBER: _ClassVar[int]
    ITEMS_FIELD_NUMBER: _ClassVar[int]
    unit: str
    delimiter: int
    period: int
    items: _containers.RepeatedCompositeFieldContainer[TimeSeriesItem]
    def __init__(self, unit: _Optional[str] = ..., delimiter: _Optional[int] = ..., period: _Optional[int] = ..., items: _Optional[_Iterable[_Union[TimeSeriesItem, _Mapping]]] = ...) -> None: ...

class GetPredictsRequest(_message.Message):
    __slots__ = ("username",)
    USERNAME_FIELD_NUMBER: _ClassVar[int]
    username: str
    def __init__(self, username: _Optional[str] = ...) -> None: ...

class GetPredictsResponse(_message.Message):
    __slots__ = ("names",)
    NAMES_FIELD_NUMBER: _ClassVar[int]
    names: _containers.RepeatedScalarFieldContainer[str]
    def __init__(self, names: _Optional[_Iterable[str]] = ...) -> None: ...
