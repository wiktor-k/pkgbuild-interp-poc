# Maintainer: Christian Rebischke <chris.rebischke@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Maxim Baz <archlinux at maximbaz dot com>
# Contributor: Daniel Martí <mvdan@mvdan.cc>

pkgname=shfmt
pkgver=3.14.1
pkgrel=1
pkgdesc="Format shell programs"
url="https://github.com/mvdan/sh"
license=(BSD-3-Clause)
arch=(x86_64)
depends=(glibc)
makedepends=(
  git
  go
  scdoc
)
options=(!lto)
source=("git+$url.git#tag=v$pkgver")
b2sums=('3997992424addca49f90f0f0be5858b233348f118f8b13f5f10da60f07de133fa1ce0bb8a6172348ff5a7d09b797bd623aa00b3c87898779151df76e1830ba57')

prepare() {
  cd "sh"
  GOFLAGS="-mod=readonly" go mod vendor -v
}

build() {
  cd "sh"
  export CGO_CPPFLAGS="$CPPFLAGS"
  export CGO_CFLAGS="$CFLAGS"
  export CGO_CXXFLAGS="$CXXFLAGS"
  export CGO_LDFLAGS="$LDFLAGS"
  export GOFLAGS="-buildmode=pie -mod=vendor -modcacherw"
  export GOPATH="$srcdir"

  local ld_flags="-compressdwarf=false -linkmode=external"
  go build -v \
    -ldflags "$ld_flags" \
    ./cmd/shfmt/...

  scdoc < cmd/shfmt/shfmt.1.scd > shfmt.1
}

check() {
  cd "sh"
  # Ensure the executable reports the correct version.
  test "$(./shfmt --version)" = "v$pkgver"

  go test ./cmd/shfmt/...
}

package() {
  cd "sh"
  install -vDm755 -t "$pkgdir/usr/bin" shfmt
  install -vDm644 -t "$pkgdir/usr/share/man/man1" shfmt.1
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
