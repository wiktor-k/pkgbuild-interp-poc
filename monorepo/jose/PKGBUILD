# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Jonas Witschel <diabonas@archlinux.org>

pkgname=jose
pkgver=15
pkgrel=2
pkgdesc='C-language implementation of Javascript Object Signing and Encryption'
arch=(x86_64)
url='https://github.com/latchset/jose'
license=(Apache-2.0)
depends=(
  glibc
  libgcc
  jansson
  openssl
  zlib
)
makedepends=(
  git
  asciidoc
  meson
)
source=("$pkgname::git+$url#tag=v$pkgver")
sha512sums=('32ab77dd54930dfbc228d44bdc303ae0be7fecec1d9c508ba4db936a55d83f8ab62c8256d1588118992857d75a7519e475a632ab99c1ca2c89b9e774ed6a8ab1')
b2sums=('37cf2c02c96cf941506ad816373f2e7dfa64e1e2cc3ed785d2f2031716bbb1253974d9b671783feb4b087af5953259cc82aa64a4480e75cc3285dbd2693fcdfd')

build() {
	cd "$pkgname"

	meson --prefix=/usr --buildtype=plain build
	meson compile -C build
}

check() {
	cd "$pkgname"

	meson test -C build
}

package() {
	cd "$pkgname"

	DESTDIR="$pkgdir" meson install -C build
}
