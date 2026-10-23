# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=plasma-pass
pkgver=1.3.1
pkgrel=2
arch=(x86_64)
pkgdesc='Plasma applet for the Pass password manager'
url='https://www.dvratil.cz/2018/05/plasma-pass/'
license=(GPL-2.0-or-later)
depends=(glibc
         gpgmepp
         kcoreaddons
         ki18n
         kirigami
         kitemmodels
         libplasma
         libstdc++
         oath-toolkit
         plasma-workspace
         plasma5support
         qgpgme
         qt6-base
         qt6-declarative)
makedepends=(extra-cmake-modules)
source=(https://download.kde.org/stable/$pkgname/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('5a88985dc522ca3287f13ed989609a93871996d4e8c5b0a41594e2d3a2fe8f2a'
            'SKIP')
validpgpkeys=('0ABDFA55A4E6BEA99A83EA974D69557AECB13683'  # Daniel Vrátil <dvratil@kde.org>
              'D81C0CB38EB725EF6691C385BB463350D6EF31EF') # Heiko Becker <heiko.becker@kde.org>

build() {
  cmake -B build -S $pkgname-$pkgver
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
