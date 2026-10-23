# Maintainer: BlackIkeEagle < ike DOT devolder AT gmail DOT com >
# Contributor: Maxim Baz <$pkgname at maximbaz dot com>
# Contributor: Batou <batou at cryptolab net>
# Contributor: Karol "Kenji Takahashi" Woźniak <kenji.sx>

pkgname=copyq
pkgver=16.0.0
pkgrel=1
pkgdesc="Clipboard manager with searchable and editable history"
url="https://github.com/hluk/${pkgname}"
depends=('hicolor-icon-theme' 'qt6-svg' 'knotifications' 'kstatusnotifieritem' 'kguiaddons' 'libxtst' 'qca-qt6' 'qtkeychain-qt6' 'miniaudio')
makedepends=('extra-cmake-modules' 'qt6-tools' 'vulkan-headers')
license=('GPL-3.0-or-later')
arch=('x86_64')
source=("${pkgname}-${pkgver}.tar.gz::${url}/archive/v${pkgver}.tar.gz")
sha256sums=('3d876f74cf32913e37b1cdeaca0c61cce857174f064bb128f76fd0fa3abdb326')

#prepare() {
    #cd CopyQ-$pkgver
    #patch -p1 -i "$srcdir/3268.patch"
#}

build() {
    cmake -B build -S CopyQ-$pkgver \
      -DMINIAUDIO_INCLUDE_DIR=/usr/include/miniaudio \
      -DCMAKE_INSTALL_PREFIX=/usr \
      -DWITH_QT6=ON
    cmake --build build
}

package() {
    DESTDIR="${pkgdir}" cmake --install build
}
