# Maintainer: Antonio Rojas <arojas@archlinux.org>
# Contributor: Fabio 'Lolix' Loli <lolix@disroot.org> -> https://github.com/FabioLolix

pkgname=strawberry
pkgver=1.2.31
pkgrel=1
pkgdesc='A music player aimed at audio enthusiasts and music collectors'
arch=(x86_64)
url='https://www.strawberrymusicplayer.org/'
license=(GPL-3.0-or-later)
depends=(alsa-lib
         chromaprint
         fftw
         gdk-pixbuf2
         glib2
         glibc
         gst-plugins-base
         gst-plugins-base-libs
         gst-plugins-good 
         gstreamer
         icu
         kdsingleapplication
         libcdio
         libebur128
         libgcc
         libgpod
         libmtp
         libpulse
         libsecret
         libstdc++
         libx11
         openssl
         qt6-base
         sqlite
         taglib
         uchardet
         udisks2)
makedepends=(boost
             cmake
             git
             qt6-tools
             rapidjson
             sparsehash)
optdepends=('gst-libav: additional codecs'
            'gst-plugins-bad: additional codecs'
            'gst-plugins-ugly: additional codecs')
source=(git+https://github.com/strawberrymusicplayer/strawberry#tag=$pkgver
        https://github.com/strawberrymusicplayer/strawberry/releases/download/$pkgver/$pkgname-$pkgver.tar.xz)
sha256sums=('19cf94a546c30e84783222c67e412049ba6bea6baa6d0ed00e94c3f0cac099fd'
            'cffc0cd453f3126a930b9e67f547e54d9fde0c8754fdead8185d2b2f5c68c037')

prepare() {
# Get API credentials from release tarball
  cp $pkgname-$pkgver/src/apicredentials.h $pkgname/src
}

build() {
  cmake -B build -S $pkgname \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -DCMAKE_CXX_FLAGS="$CXXFLAGS -DQT_NO_DEBUG_OUTPUT"
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
