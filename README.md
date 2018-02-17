# go-marc21-details

[![GoDoc](https://godoc.org/github.com/gsiems/go-marc21/pkg/marc21-details?status.svg)](https://godoc.org/github.com/gsiems/go-marc21-details/pkg/details)
[![Go Report Card](https://goreportcard.com/badge/github.com/gsiems/go-marc21-details)](https://goreportcard.com/report/github.com/gsiems/go-marc21-details)

Extract detailed information from MARC21 records.

The intent is to parse a MARC record and translate the values of the
different elements into human readable output so that one does not need
to be a MARC21 expert to determine what information is recorded in a
MARC record.

Currently parses the leader and control fields for a MARC record.
