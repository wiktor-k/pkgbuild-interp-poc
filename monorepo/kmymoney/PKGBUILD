# Maintainer: Jaroslav Lichtblau <svetlemodry@archlinux.org>
# Maintainer: Carl Smedstad <carsme@archlinux.org>
# Contributor: Laurent Carlier <lordheavym@gmail.com>
# Contributor: Vamp898 <vamp898@web.de>
# Contributor: Jaroslaw Swierczynski <swiergot@aur.archlinux.org>
# Contributor: Eric Belanger <eric@archlinux.org>
# Contributor: Todd Maynard <arch@toddmaynard.com>

pkgname=kmymoney
pkgver=5.2.2
pkgrel=8
pkgdesc="Personal finance manager for KDE which operates similarly to MS-Money or Quicken"
arch=('x86_64')
url="https://kmymoney.org/"
license=('GPL-2.0-or-later')
depends=(
  'akonadi'
  'aqbanking'
  'glibc'
  'gmp'
  'gpgmepp'
  'gwenhywfar'
  'gwenhywfar-qt'
  'hicolor-icon-theme'
  'karchive'
  'kcmutils'
  'kcolorscheme'
  'kcompletion'
  'kconfig'
  'kconfigwidgets'
  'kcontacts'
  'kcoreaddons'
  'kcrash'
  'kdiagram'
  'kholidays'
  'ki18n'
  'kidentitymanagement'
  'kio'
  'kitemmodels'
  'kitemviews'
  'kjobwidgets'
  'ktextwidgets'
  'kwidgetsaddons'
  'kxmlgui'
  'libalkimia'
  'libgcc'
  'libical'
  'libofx'
  'libstdc++'
  'python'
  'qt6-base'
  'qt6-declarative'
  'qtkeychain-qt6'
  'sonnet'
  'sqlcipher'
)
makedepends=(
  'extra-cmake-modules'
  'kdoctools'
  'qgpgme'
)
optdepends=('perl: for financequote.pl')
source=("https://download.kde.org/stable/$pkgname/$pkgver/$pkgname-$pkgver.tar.xz"{,.sig})
b2sums=('701ef89f27c50d3b297ff59c8a4b30e6852786fdadabef2d08e4fa17740dbdeb5019a1851a852748fe7808eb29b8aa57afd20f035067dfb66f31612e9be83352'
        'SKIP')
validpgpkeys=('D69A745A55331F44F404D8258D4DE062AA2EB01C') # Thomas Baumgart <thb@net-bembel.de>

build() {
  cmake -B build -S $pkgname-$pkgver \
    -DCMAKE_BUILD_TYPE=None \
    -DCMAKE_INSTALL_PREFIX=/usr \
    -Wno-dev \
    -DBUILD_WITH_QT6=ON \
    -DBUILD_TESTING=ON
  cmake --build build
}

check() {
  ctest --test-dir=build --output-on-failure
}

package() {
  DESTDIR="$pkgdir" cmake --install build
}
