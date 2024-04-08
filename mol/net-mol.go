/*
 * Copyright (c) 2024.
 * Created by Andy Pangaribuan <https://github.com/apangaribuan>.
 *
 * This product is protected by copyright and distributed under
 * licenses restricting copying, distribution and decompilation.
 * All Rights Reserved.
 */

package mol

import "crypto/tls"

type NetOpt struct {
	UsingClientLB         *bool
	TlsConfig             *tls.Config
	TlsCertFile           *string
	TlsServerNameOverride *string
}
