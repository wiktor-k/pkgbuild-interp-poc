# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: Matteo Giordano <mail@matteogiordano.me>
pkgname=superfile
pkgver=1.6.0
pkgrel=1
pkgdesc='Pretty fancy and modern terminal file manager'
arch=('x86_64')
url='https://superfile.netlify.app/'
license=('MIT')
depends=('glibc')
makedepends=('go')
optdepends=(
  'exiftool: metadata plugin'
)
source=("https://github.com/yorukot/superfile/archive/refs/tags/v${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha256sums=('7340ce9e3e7db401164310dd5d2d1dfa3ccf76118ac49d192b0fac2a292c5b0d')

build() {
    cd "${pkgname}-${pkgver}"
    export CGO_CPPFLAGS="${CPPFLAGS}"
    export CGO_CFLAGS="${CFLAGS}"
    export CGO_CXXFLAGS="${CXXFLAGS}"
    export CGO_LDFLAGS="${LDFLAGS}"
    export GOFLAGS="-buildmode=pie -trimpath -ldflags=-linkmode=external -mod=readonly -modcacherw"
    go build -o "./bin/spf"
}

package() {
    cd "${pkgname}-${pkgver}"
    install -Dm755 bin/spf -t "${pkgdir}/usr/bin/"
    install -Dm644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}/"
    install -Dm644 README.md -t "${pkgdir}/usr/share/doc/${pkgname}/"
}
