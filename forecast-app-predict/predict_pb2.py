# -*- coding: utf-8 -*-
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: predict.proto
# Protobuf Python Version: 4.25.1
"""Generated protocol buffer code."""
from google.protobuf import descriptor as _descriptor
from google.protobuf import descriptor_pool as _descriptor_pool
from google.protobuf import symbol_database as _symbol_database
from google.protobuf.internal import builder as _builder
# @@protoc_insertion_point(imports)

_sym_db = _symbol_database.Default()


from google.protobuf import timestamp_pb2 as google_dot_protobuf_dot_timestamp__pb2


DESCRIPTOR = _descriptor_pool.Default().AddSerializedFile(b'\n\rpredict.proto\x1a\x1fgoogle/protobuf/timestamp.proto\"G\n\x0eTimeSeriesItem\x12&\n\x02ts\x18\x01 \x01(\x0b\x32\x1a.google.protobuf.Timestamp\x12\r\n\x05value\x18\x02 \x01(\x01\"b\n\x12MakePredictRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\x12\x0c\n\x04unit\x18\x03 \x01(\t\x12\x1e\n\x05items\x18\x04 \x03(\x0b\x32\x0f.TimeSeriesItem\"5\n\x13MakePredictResponse\x12\x1e\n\x05items\x18\x01 \x03(\x0b\x32\x0f.TimeSeriesItem\"3\n\x11GetPredictRequest\x12\x10\n\x08username\x18\x01 \x01(\t\x12\x0c\n\x04name\x18\x02 \x01(\t\"\"\n\x12GetPredictResponse\x12\x0c\n\x04\x64\x61ta\x18\x01 \x01(\x0c\x32\x81\x01\n\x0ePredictService\x12\x38\n\x0bMakePredict\x12\x13.MakePredictRequest\x1a\x14.MakePredictResponse\x12\x35\n\nGetPredict\x12\x12.GetPredictRequest\x1a\x13.GetPredictResponseB\x0cZ\ngo/predictb\x06proto3')

_globals = globals()
_builder.BuildMessageAndEnumDescriptors(DESCRIPTOR, _globals)
_builder.BuildTopDescriptorsAndMessages(DESCRIPTOR, 'predict_pb2', _globals)
if _descriptor._USE_C_DESCRIPTORS == False:
  _globals['DESCRIPTOR']._options = None
  _globals['DESCRIPTOR']._serialized_options = b'Z\ngo/predict'
  _globals['_TIMESERIESITEM']._serialized_start=50
  _globals['_TIMESERIESITEM']._serialized_end=121
  _globals['_MAKEPREDICTREQUEST']._serialized_start=123
  _globals['_MAKEPREDICTREQUEST']._serialized_end=221
  _globals['_MAKEPREDICTRESPONSE']._serialized_start=223
  _globals['_MAKEPREDICTRESPONSE']._serialized_end=276
  _globals['_GETPREDICTREQUEST']._serialized_start=278
  _globals['_GETPREDICTREQUEST']._serialized_end=329
  _globals['_GETPREDICTRESPONSE']._serialized_start=331
  _globals['_GETPREDICTRESPONSE']._serialized_end=365
  _globals['_PREDICTSERVICE']._serialized_start=368
  _globals['_PREDICTSERVICE']._serialized_end=497
# @@protoc_insertion_point(module_scope)
