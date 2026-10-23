# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=skanpage
pkgver=26.08.2
pkgrel=2
pkgdesc='Utility to scan images and multi-page documents'
arch=(x86_64)
url='https://apps.kde.org/skanpage/'
license=(GPL-2.0-or-later)
depends=(glibc
         kconfig
         kcoreaddons
         kcrash
         ki18n
         kio
         kirigami
         kquickimageeditor
         ksanecore
         kwidgetsaddons
         kxmlgui
         leptonica
         libstdc++
         purpose
         qt6-base
         qt6-declarative
         qt6-webengine
         tesseract)
makedepends=(extra-cmake-modules)
groups=(kde-applications
        kde-utilities)
source=(https://download.kde.org/stable/release-service/$pkgver/src/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('1e06c842ee774de21a02772c5bb646088c6de970af22e9b3e879273b615bfe6e'
            'SKIP')
validpgpkeys=(CA262C6C83DE4D2FB28A332A3A6A4DB839EAA6D7  # Albert Astals Cid <aacid@kde.org>
              F23275E4BF10AFC1DF6914A6DBD2CE893E2D1C87  # Christoph Feck <cfeck@kde.org>
              D81C0CB38EB725EF6691C385BB463350D6EF31EF) # Heiko Becker <heiko.becker@kde.org>

build() {
  cmake -B build -S $pkgname-$pkgver \
    -DBUILD_TESTING=OFF
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
