# Maintainer: Giancarlo Razzolini <grazzolini@archlinux.org>
# Maintainer: Christian Heusel <gromit@archlinux.org>
# Contributor: Guillaume ALAUX <guillaume@archlinux.org>
# Contributor: Allan McRae <allan@archlinux.org>
# Contributor: fancris3 <fancris3 at gmail.com>
# Contributor: Daniel J Griffiths <ghost1227@archlinux.us>

pkgname=terminator
pkgver=2.1.6
pkgrel=1
pkgdesc='Terminal emulator that supports tabs and grids'
arch=('any')
url='https://github.com/gnome-terminator/terminator'
license=('GPL-2.0-only')
depends=(
    'gsettings-desktop-schemas'
    'libkeybinder3'
    'libnotify'
    'python'
    'python-cairo'
    'python-configobj'
    'python-dbus'
    'python-gobject'
    'python-psutil'
    'vte3'
    'xdg-utils'
)
makedepends=(
    'intltool'
    'python-build'
    'python-installer'
    'python-setuptools'
    'python-wheel'
)
source=(${url}/releases/download/v${pkgver}/${pkgname}-${pkgver}.tar.gz{,.asc}
        0001-fixed-validate-module-import-for-newer-version-of-co.patch)

validpgpkeys=('6FBDB41587919A0F4BA96EBA9A358BAB04A4ED27'  #Markus Frosch <markus@lazyfrosch.de>
              'D11A7596F61705480C711598F2FAC7C7BAE930A5') #Matt Rose <mattrose@folkwolf.net>
sha512sums=('1fe350056c85b697af06681d2957a4d4e6dbd668240d10c0b73ad95b4877786c22f11d7ab97230776b729c487c1cfeb196b03d973e4fa82e7b4a75e960d22ac5'
            'SKIP'
            'da609bbb1f2fce64bd345ec43c05f6b8d76046cfbd1be764d6fcd555d306dfca1dd5786b0fc65c0984e49a85454c9eeabf5a1f8c960412c14a97c5350912f8a9')
b2sums=('634ba872a4dd4acdc84401095156e6d6d4b07efd0037ffcffac592f905e27989b637c8720f1827b4927736941e00ee3d18f7eded20e8fa673206e4832c05d52c'
        'SKIP'
        '7cde90cefde34346ef20b50ebd5b930cb47df650c649e199adbb1b6235db0b40c0c2ad2b0e30bc764d52469357ce0b0424648f70da93b10d39ed0793cbbfe648')

prepare () {
    cd "${pkgname}-${pkgver}"

    # Fix for python configobj validate
    patch -Np1 < $srcdir/0001-fixed-validate-module-import-for-newer-version-of-co.patch

    # Remove pytest-runner from setup_requires
    sed -i '/pytest-runner/d' setup.py
}

build() {
    cd "${pkgname}-${pkgver}"

    python -m build --wheel --no-isolation
}

package() {
    cd "${pkgname}-${pkgver}"

    python -m installer --destdir="$pkgdir" dist/*.whl

    rm -f "${pkgdir}/usr/share/icons/hicolor/icon-theme.cache"
}
