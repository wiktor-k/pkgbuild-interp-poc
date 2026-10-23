# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Maintainer: Christian Heusel <gromit@archlinux.org>
# Contributor: hexchain <arch at hexchain.org>

pkgname=libkrun
pkgver=1.19.7
pkgrel=1
pkgdesc="A dynamic library providing Virtualization-based process isolation capabilities"
url='https://github.com/containers/libkrun'
arch=('x86_64')
license=('Apache-2.0')
makedepends=('cargo' 'patchelf' 'clang')
depends=('glibc' 'libgcc' 'zstd' 'bzip2' 'libkrunfw' 'libpipewire' 'virglrenderer')
source=("https://github.com/containers/libkrun/archive/refs/tags/v$pkgver/$pkgname-$pkgver.tar.gz")
sha256sums=('d4541fcd959ad43fb073bff61bc91b645ea6413a10b8319e9b61664a6102d20c')

prepare() {
  cd "$pkgname-$pkgver"

  cargo fetch --locked --target "$(rustc --print host-tuple)"
}

build() {
  cd "$pkgname-$pkgver"

  export ZSTD_SYS_USE_PKG_CONFIG=1
  make BLK=1 NET=1 EFI=0 GPU=1 SND=1 INPUT=1 VERBOSE=1
}

package() {
  cd "$pkgname-$pkgver"

  make DESTDIR="$pkgdir" PREFIX=/usr LIBDIR_Linux=lib install

  install -Dm644 LICENSE "$pkgdir"/usr/share/licenses/$pkgname/LICENSE
}
