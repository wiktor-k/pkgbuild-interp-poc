# Maintainer: Bruno Pagani <archange@archlinux.org>
# Maintainer: Caleb Maclennan <caleb@alerque.com>
# Contributor: Gaetan Bisson <bisson@archlinux.org>
# Contributor: Tobias Kieslich <tobias@archlinux.org>
# Contributor: Giovanni Scafora <giovanni@archlinux.org>
# Contributor: Dominik Ryba <domryba@post.pl>

pkgname=hugin
pkgver=2025.0.1
pkgrel=5
pkgdesc="Panorama photo stitcher"
arch=(x86_64)
url="http://hugin.sourceforge.net/"
license=(GPL-2.0-only)
depends=(enblend-enfuse
         exiv2
         fftw
         glu
         lapack
         lcms2
         lensfun
         libepoxy
         libpano13
         libtiff
         libxi
         libxmu
         openexr
         python
         sqlite
         vigra
         wxwidgets-gtk3)
makedepends=(cmake
             dos2unix
             mesa
             python-setuptools
             swig
             tclap)
optdepends=('darktable: RAW import using darktable'
            'dcraw: RAW import using dcraw'
            'perl-image-exiftool: GPano tags support'
            'rawtherapee: RAW import using rawtherapee')
_archive="$pkgname-$pkgver"
source=("https://downloads.sourceforge.net/$pkgname/$_archive.tar.bz2"
        ignore_gzip_timestamps.patch
        fix-vigra-version-check.patch
        swig-4.5.patch)
sha256sums=('7cf8eb33a6a8848cc7f816faf4bc88389228883d5513136dccb5cb243912ab79'
            'e6a41ce195c55f0c895a359203c84c47e885dbbd98c9b077c9a5e092be3bafb4'
            '3341a9f8cb5635c9d1b169317a67f5386aeeb09573dd019523cd15cc698ca3a5'
            '8e58e7ba909e80f6142d1e44f1be559817e976f2f1f1bd51db18705fb98dff43')

prepare() {
# Ignore timestamps recording in gzip metadata
# This is required for reproducible builds
  cd "$_archive"
  patch -Np1 -i "$srcdir/ignore_gzip_timestamps.patch"

  patch -p1 -i ../fix-vigra-version-check.patch
  patch -p1 -i ../swig-4.5.patch
}

build() {
  cmake -B build -S "$_archive" \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DENABLE_LAPACK=yes \
    -DBUILD_HSI=ON \
    -DUSE_GDKBACKEND_X11=ON \
    -DBUILD_WITH_EPOXY=ON
  make -C build
}

package() {
  make -C build DESTDIR="$pkgdir" install
}
