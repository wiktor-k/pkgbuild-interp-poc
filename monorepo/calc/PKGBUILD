# Maintainer: Andrew Crerar <crerar@archlinux.org>
# Contributor: Kyle Keen <keenerd@gmail.com>
# Contributor: Sergej Pupykin <pupykin.s+arch@gmail.com>

pkgname=calc
pkgver=2.17.0.2
pkgrel=1
pkgdesc="Arbitrary precision console calculator"
arch=(x86_64)
url="https://github.com/lcn2/calc"
license=(LGPL-2.1-only)
depends=(readline)
makedepends=(mandoc)
source=(${pkgname}-${pkgver}.tar.gz::https://github.com/lcn2/${pkgname}/archive/v${pkgver}.tar.gz)
b2sums=('5db4c1885e2568ae68517c9f7b07c81c1df626fac58da022c69d25e4e4effa824a3fc69ffee3e9ff1cfe7d9f98b74a4979b96283b6077c2eae3c865f0fe4c15f')

prepare() {
  cd "${pkgname}-${pkgver}"

  sed -i 's/${CC} ${LIBCALC_SHLIB}/${CC} ${LDFLAGS} ${LIBCALC_SHLIB}/' Makefile
  sed -i 's/${CC} ${LIBCUSTCALC_SHLIB}/${CC} ${LDFLAGS} ${LIBCUSTCALC_SHLIB}/' custom/Makefile
}

build() {
  cd "${pkgname}-${pkgver}"

  make \
    USE_READLINE="-DUSE_READLINE" \
    READLINE_LIB="-lreadline" \
    READLINE_EXTRAS="-lhistory -lncurses" \
    EXTRA_CFLAGS="${CPPFLAGS} ${CFLAGS} -Wno-error=format-security" \
    EXTRA_LDFLAGS="${LDFLAGS}" \
    ARCH_CFLAGS="" \
    LD_SHARE="" \
    DEBUG=""
}

check() {
  cd "${pkgname}-${pkgver}"

  make chk
}

package() {
  cd "${pkgname}-${pkgver}"

  make \
    T="${pkgdir}" install

  # `cscript` is a directory of example calc scripts, not a binary. Move it to
  # the correct location.
  mv "${pkgdir}/usr/bin/cscript" "${pkgdir}/usr/share/calc/"
}
