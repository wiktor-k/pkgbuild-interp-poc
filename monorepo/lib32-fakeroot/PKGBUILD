# Maintainer: Felix Yan <felixonmars@archlinux.org>
# Contributor: Jan Alexander Steffens (heftig) <jan.steffens@gmail.com>
# Contributor: Allan McRae <allan@archlinux.org>
# Contributor: Jochem Kossen <j.kossen@home.nl>

_pkgbasename=fakeroot
pkgname=lib32-${_pkgbasename}
epoch=1
pkgver=1.37.2
pkgrel=3
pkgdesc='Tool for simulating superuser privileges (32-bit)'
arch=('x86_64')
license=('GPL-3.0-or-later')
url="https://tracker.debian.org/pkg/fakeroot"
depends=('lib32-glibc' "$_pkgbasename")
makedepends=('git' 'po4a')
# Build from my mirror with original commit and historical tags
# unter upstream manages to do a proper release...
source=("fakeroot-archive::git+https://github.com/eworm-de/fakeroot.git#tag=upstream/${pkgver}-2026-01-17-8d7bc2eb")
sha256sums=('46f44681cc939dd3f127761717830c0bfa72e1568191d4df071ae297013591e3')

prepare() {
  ln -s fakeroot-archive fakeroot
  cd "${_pkgbasename}"

  autoreconf -fi
}

build() {
  export CC="gcc -m32"
  export CXX="g++ -m32"
  export PKG_CONFIG_PATH="/usr/lib32/pkgconfig"

  cd "${_pkgbasename}"

  ./configure \
    --prefix=/usr \
    --libdir=/usr/lib32/libfakeroot \
    --disable-static \
    --with-ipc=sysv
  make

  cd doc
  po4a -k 0 --rm-backups --variable "srcdir=../doc/" po4a/po4a.cfg
}

package() {
  cd "${_pkgbasename}"

  make DESTDIR="${pkgdir}" install
  rm -rf "${pkgdir}"/usr/{bin,share}

  install -dm0755 "${pkgdir}/etc/ld.so.conf.d/"
  echo '/usr/lib32/libfakeroot' > "${pkgdir}/etc/ld.so.conf.d/${pkgname}.conf"
}
