// This file was generated by nomdl/codegen.
// @flow
/* eslint-disable */

import {
  boolType as _boolType,
  makeSetType as _makeSetType,
  newSet as _newSet,
} from '@attic/noms';
import type {
  NomsSet as _NomsSet,
} from '@attic/noms';


export function newSetOfBool(values: Array<boolean>): Promise<_NomsSet<boolean>> {
  return _newSet(values, _makeSetType(_boolType));
}