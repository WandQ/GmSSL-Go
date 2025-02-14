/*
 *  Copyright 2014-2023 The GmSSL Project. All Rights Reserved.
 *
 *  Licensed under the Apache License, Version 2.0 (the License); you may
 *  not use this file except in compliance with the License.
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 */
/* +build cgo */
package gmssl3

/*
#include <gmssl/version.h>
*/
import "C"

const (
	GmSSLGoVersion = "1.3.1"
)

func GetGmSSLLibraryVersion() string {
	return C.GoString(C.gmssl_version_str())
}
