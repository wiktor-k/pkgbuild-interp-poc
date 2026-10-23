# Maintainer: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Martin Wimpress <code@flexion.org>
# Contributor: Balló György <ballogyor+arch at gmail dot com>
# Contributor: Bartłomiej Piotrowski <barthalion@gmail.com>
# Contributor: SpepS <dreamspepser at yahoo dot it>
# Contributor: Alexander Fehr <pizzapunk gmail com>

pkgname=galculator
pkgver=2.1.4
pkgrel=12
pkgdesc='GTK based scientific calculator'
arch=(x86_64)
url='https://github.com/galculator/galculator'
license=(GPL-2.0-or-later)
depends=(
  glib2
  glibc
  gtk3
  hicolor-icon-theme
  libquadmath
  pango
)
makedepends=(
  git
  intltool
)
source=(
  "git+https://github.com/galculator/galculator.git#tag=v$pkgver"
  0001-Fix-multiple-definition-of-prefs-compile-error-with-.patch
  0002-Declare-function-parameters-as-required-by-C23.patch
  galculator-sign-accel.patch
)
b2sums=(
  4a204b3d459f0f9e9e7014e39d1420f975f3569e8f36acefd5b5ceb9bcec153435e5f7dd40aa9d678f98a54e979739b44494bad1f8910b645719c4946bfe6bbf
  9ab69d9ab67edde7a2cc962013a7156194d822601429a7f62066c02a078c7fdc8d0811ca6400e72e35184e83c941a6c270c6976fd5cf75a7800194bd123b6c9d
  8085ef6f869396f29d8385eccb2d14ff58ca99ab2a67de60ded3ba60a6ed0ad5601a76cc46528f617e9c97d52920e297eb777711bdff92c4a6a0866ece6e766b
  0624ea8ea25432b9eef08ada9a88f3ca17695a72e7c14ef85ed0feaec7b5f09d76f58f4bf7993caf518061fdc22cc2bb2772ca7559489df1682d75315e299c87
)

prepare() {
  cd $pkgname

  # Fix build with recent GCC versions
  git apply -3 ../0001-Fix-multiple-definition-of-prefs-compile-error-with-.patch
  git apply -3 ../0002-Declare-function-parameters-as-required-by-C23.patch

  # Add keyboard accelerators for change sign button
  # https://github.com/galculator/galculator/pull/75
  git apply -3 ../galculator-sign-accel.patch
  
  sed -i 's|$(datadir)/appdata|$(datadir)/metainfo|' Makefile.am
  sed -i 's/galculator.xpm//' pixmaps/Makefile.am
  autoreconf -fiv
}

build() {
  cd $pkgname
  ./configure \
    --prefix=/usr \
    --sysconfdir=/etc \
    --localstatedir=/var
  make
}

package() {
  cd $pkgname
  make DESTDIR="$pkgdir" install
}
