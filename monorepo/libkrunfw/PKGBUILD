# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: hexchain <arch at hexchain.org>
pkgname=libkrunfw
# Check https://github.com/containers/libkrunfw/blob/main/Makefile#L1 for kernel version
_kver=6.12.68
pkgver=5.6.2
pkgrel=1
pkgdesc="A dynamic library bundling the guest payload consumed by libkrun"
url="https://github.com/containers/libkrunfw"
arch=('x86_64')
license=('GPL-2.0-only' 'LGPL-2.1-only')
depends=('glibc')
makedepends=('bc' 'python-pyelftools' 'cpio')
source=("https://github.com/containers/libkrunfw/archive/refs/tags/v$pkgver/$pkgname-$pkgver.tar.gz"
        "https://cdn.kernel.org/pub/linux/kernel/v${_kver%%.*}.x/linux-${_kver}.tar.xz")
noextract=("linux-${_kver}.tar.xz")
options=(!debug)  # suppress empty debug package
sha256sums=('df45d649fcbd07a4d0ca03fa836b2f640fdcf9b40187f3eb7023259c6e83d582'
            'd367c7504bd4da520dd01eb08125d2d0ac088bc8af4cd56d236f2074dd4225b7')

build() {
  cd "$pkgname-$pkgver"

  mkdir tarballs
  cp -s "$srcdir/linux-${_kver}.tar.xz" tarballs/

  make
}

package() {
  cd "$pkgname-$pkgver"

  make DESTDIR="$pkgdir" PREFIX=/usr LIBDIR_Linux=lib install

  install -Dm644 LICENSE-GPL-2.0-only "$pkgdir"/usr/share/licenses/$pkgname/LICENSE-GPL-2.0-only
  install -Dm644 LICENSE-LGPL-2.1-only "$pkgdir"/usr/share/licenses/$pkgname/LICENSE-LGPL-2.1-only
}
