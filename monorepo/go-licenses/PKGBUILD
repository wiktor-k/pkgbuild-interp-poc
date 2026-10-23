# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>

pkgname=go-licenses
pkgver=2.0.1
pkgrel=1
pkgdesc="Reports on the licenses used by a Go package and its dependencies"
url="https://github.com/google/go-licenses"
arch=('x86_64')
license=('Apache-2.0')
depends=(
  'glibc'
)
makedepends=(
  'go'
)
# we can not use LTO as otherwise we get no reproducible package with full RELRO
options=(!lto)
source=(${pkgname}-${pkgver}.tar.gz::https://github.com/google/${pkgname}/archive/refs/tags/v${pkgver}.tar.gz)
sha256sums=('04a875604c0cc62217fea3d283c243a2e4298142cdfeafb088b9e486daabded3')
b2sums=('b726b41d87110bb0f04b4a1c908a01c564d5eabdb3781fe5aeea2627ca1a4e50d2184a54b8ea38033a6895ef357729177120a0601bb8b1192af1387cc2cc863f')

prepare() {
  cd ${pkgname}-${pkgver}
  go mod download
}

build() {
  cd ${pkgname}-${pkgver}
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export CGO_REQUIRED="1"

  go build \
    -buildmode=pie \
    -mod=readonly \
    -modcacherw \
    -ldflags '-compressdwarf=false -linkmode=external' \
    .
}

package() {
  cd ${pkgname}-${pkgver}
  install -Dm755 ${pkgname} "${pkgdir}/usr/bin/${pkgname}"
  install -Dm644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}

# vim: ts=2 sw=2 et:
