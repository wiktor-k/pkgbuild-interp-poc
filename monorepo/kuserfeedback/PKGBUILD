# Maintainer: Antonio Rojas <arojas@archlinux.org>

pkgname=kuserfeedback
pkgver=6.31.0
pkgrel=1
pkgdesc='Framework for collecting user feedback for applications via telemetry and surveys'
arch=(x86_64)
url='https://develop.kde.org/products/frameworks/'
license=(LGPL-2.0-only LGPL-3.0-only)
depends=(glibc
         libstdc++
         qt6-base)
makedepends=(clang
             extra-cmake-modules
             qt6-charts
             qt6-declarative
             qt6-svg
             qt6-tools)
optdepends=('qt6-charts: Feedback console'
            'qt6-declarative: QML bindings'
            'qt6-svg: Feedback console')
conflicts=(kuserfeedback5)
replaces=(kuserfeedback5)
source=(https://download.kde.org/stable/frameworks/${pkgver%.*}/$pkgname-$pkgver.tar.xz{,.sig})
sha256sums=('6a761f2eb1507110107b2ddf2a723bbe6fa205eac086ce04a4272fdfabb7c31a'
            'SKIP')
validpgpkeys=(53E6B47B45CEA3E0D5B7457758D0EE648A48B3BB # David Faure <faure@kde.org>
              90A968ACA84537CC27B99EAF2C8DF587A6D4AAC1 # Nicolas Fella <nicolas.fella@kde.org>
              )

build() {
  cmake -B build -S $pkgname-$pkgver \
    -DBUILD_TESTING=OFF
  cmake --build build
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
