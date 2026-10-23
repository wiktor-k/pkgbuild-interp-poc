# Maintainer: Bruno Pagani <archange@archlinux.org>

_pkg=filter-rspamd
pkgname=opensmtpd-${_pkg}
pkgver=0.1.9
pkgrel=1
pkgdesc="OpenSMTPD filter integration for Rspamd"
arch=(x86_64)
url=https://github.com/poolpOrg/filter-rspamd
license=(ISC)
depends=(glibc rspamd)
makedepends=(go)
source=(${url}/archive/v${pkgver}/${pkgname}-${pkgver}.tar.gz)
sha512sums=('14bcf5928fbb650c0ed54312f660b866f4fdc7f72a91bfb3879990dfc207eb0d890db9c242b58039620793d534f89781e68824bfedb78a488d426180a93db41a')

build() {
  cd ${_pkg}-${pkgver}
  export CGO_CPPFLAGS="${CPPFLAGS}"
  export CGO_CFLAGS="${CFLAGS}"
  export CGO_CXXFLAGS="${CXXFLAGS}"
  export CGO_LDFLAGS="${LDFLAGS}"
  export GOFLAGS="-buildmode=pie -trimpath -mod=readonly -modcacherw"
  go build -v -ldflags "-linkmode external" -o ${_pkg}
}

package() {
  cd ${_pkg}-$pkgver
  install -D filter-rspamd -t "${pkgdir}"/usr/lib/smtpd/opensmtpd
  install -Dm644 LICENSE -t "${pkgdir}"/usr/share/licenses/${pkgname}/
  install -Dm644 README.md -t "${pkgdir}"/usr/share/doc/${pkgname}/
}
