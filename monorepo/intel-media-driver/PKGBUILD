# Maintainer: Bruno Pagani <archange@archlinux.org>
# Maintainer: Daniel Bermond <dbermond@archlinux.org>

pkgname=intel-media-driver
pkgver=26.3.5
pkgrel=1
pkgdesc='Intel Media Driver for VAAPI — Broadwell+ iGPUs'
arch=('x86_64')
url='https://github.com/intel/media-driver/'
license=('BSD-3-Clause' 'MIT')
depends=(
    'glibc'
    'intel-gmmlib'
    'libgcc'
    'libstdc++'
    'libva')
makedepends=(
    'cmake')
source=("https://github.com/intel/media-driver/archive/intel-media-${pkgver}/${pkgname}-${pkgver}.tar.gz")
sha256sums=('76f41c06cefc91d57e0b7d366b7a91b8420ec77939326491795f9fb26afca7e7')

build() {
    cmake -B build -S "media-driver-intel-media-${pkgver}" \
        -G 'Unix Makefiles' \
        -DCMAKE_INSTALL_PREFIX:PATH='/usr' \
        -DCMAKE_INSTALL_LIBDIR:PATH='lib' \
        -DINSTALL_DRIVER_SYSCONF:BOOL='OFF' \
        -DMEDIA_BUILD_FATAL_WARNINGS:BOOL='OFF' \
        -Wno-author
    cmake --build build
}

package() {
    DESTDIR="$pkgdir" cmake --install build
    install -D -m644 "media-driver-intel-media-${pkgver}/LICENSE.md" -t "${pkgdir}/usr/share/licenses/${pkgname}"
}
