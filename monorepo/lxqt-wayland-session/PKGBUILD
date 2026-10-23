# Maintainer: Chih-Hsuan Yen <base64_decode("eXUzYWN0eHQydHR0ZmlteEBjaHllbi5jYwo=")>

pkgname=lxqt-wayland-session
pkgver=0.4.1
pkgrel=1
pkgdesc='Files needed for the LXQt Wayland Session'
arch=('any')
url='https://github.com/lxqt/lxqt-wayland-session'
# see https://github.com/lxqt/lxqt-wayland-session/blob/master/LICENSE for more details
license=(
  'LGPL-2.1-only'       # most files
  'MIT'                 # files adapted from sway and wayfire
  'GPL-2.0-only'        # files adapted from labwc https://github.com/labwc/labwc/blob/0.8.0/meson.build#L5
  'GPL-3.0-only'        # files adapted from river (https://codeberg.org/river/river/src/tag/v0.3.5#licensing) and labwc-themes
  'GPL-3.0-or-later'    # files adopted from niri (https://github.com/YaLTeR/niri/blob/v0.1.9/Cargo.toml#L8)
  'BSD-3-Clause'        # files adapted from hyprland
  'CC-BY-SA-4.0'        # wallpapers
)
depends=(
  'sh'
  'lxqt-session'
  # Allow to use xdg-utils under LXQt, similar to https://github.com/lxqt/lxqt-session/blob/2.0.0/CHANGELOG#L27
  # startlxqtwayland sets XDG_CURRENT_DESKTOP
  'qtxdg-tools'
  # For applications that need layer-shell-qt (ex: lxqt-panel)
  'layer-shell-qt'
)
optdepends=(
  # See upstream README.md for supported compositors
  # Dependencies that are not official packages are commented out
  'labwc'
  # 'wayfire'
  'kwin'
  'hyprland'
  'niri'
  'river'
  'sway'
)
makedepends=('git' 'cmake' 'qt6-tools' 'lxqt-build-tools')
source=("https://github.com/lxqt/$pkgname/releases/download/$pkgver/$pkgname-$pkgver.tar.xz"{,.asc}
        'lxqt-wayland-session-import-environment.patch')
sha256sums=('22de8fbedc8cb584d0b7d388b53c5ff288f9ad419562602f2f1abcad34c821c5'
            'SKIP'
            '339637044f205dbd2c730953166b4ed7acca54efd2b50a1aa2fcfed8d63c3feb')
validpgpkeys=(
  "19DFDF3A579BD509DBB572D8BE793007AD22DF7E"  # https://github.com/tsujan
  "48836EB124A1EAAE695EDFEA229CA0A00A17D258"  # the key owned by standreas@riseup.net as per keys.openpgp.org. That email is used by https://github.com/stefonarch/, a long-time LXQt contributor
)
install=lxqt-wayland-session.install

prepare() {
  cd "$srcdir/$pkgname-$pkgver"

  # Set XDG_CURRENT_DESKTOP for new units started by user service manager
  # https://github.com/lxqt/lxqt-wayland-session/pull/90
  patch -Np1 -i ../lxqt-wayland-session-import-environment.patch
}

build() {
  cmake -B build -S "$srcdir/$pkgname-$pkgver" \
    -DCMAKE_INSTALL_PREFIX=/usr
  make -C build
}

package() {
  make -C build DESTDIR="$pkgdir" install

  install -Dm644 $pkgname-$pkgver/{COPYING.LESSER,COPYING,LICENSE.MIT,LICENSE.GPLv2,LICENSE.BSD} -t "$pkgdir"/usr/share/licenses/$pkgname
}
