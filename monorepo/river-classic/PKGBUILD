# Maintainer: Daurnimator <daurnimator@archlinux.org>
# Contributor: Andrea Feletto <andrea@andreafeletto.com>

pkgname=river-classic
pkgver=0.3.17
pkgrel=2
pkgdesc="A dynamic tiling wayland compositor"
arch=('x86_64')
url="https://codeberg.org/river/river-classic"
license=('GPL-3.0-or-later')
depends=('glibc'
         'libevdev'
         'libinput'
         'libxkbcommon'
         'mesa'
         'pixman'
         'sh'
         'wayland'
         'wlroots0.20'
         'xorg-xwayland')
makedepends=('git'
             'scdoc'
             'wayland-protocols'
             'zig'
             # https://codeberg.org/ziglang/zig/issues/31272
             'clang')
optdepends=('polkit: access seat through systemd-logind')
provides=('wayland-compositor')
conflicts=('river')
source=(git+https://codeberg.org/river/river-classic.git#tag=v${pkgver}?signed)
#
# disable source caching for now: https://codeberg.org/ziglang/zig/issues/31866
#
# PACKAGING.md -> build.zig.zon
# source+=(zig-pixman-v0.3.0.tar.gz::https://codeberg.org/ifreund/zig-pixman/archive/v0.3.0.tar.gz
         # zig-wayland-v0.6.0.tar.gz::https://codeberg.org/ifreund/zig-wayland/archive/v0.6.0.tar.gz
         # zig-wlroots-v0.20.1.tar.gz::https://codeberg.org/ifreund/zig-wlroots/archive/v0.20.1.tar.gz
         # zig-xkbcommon-v0.3.0.tar.gz::https://codeberg.org/ifreund/zig-xkbcommon/archive/v0.3.0.tar.gz
         # translate_c::git+https://codeberg.org/ziglang/translate-c.git#commit=7a1a9fdc4ab00835748a6657ecbb835e3d5d45f7)
noextract=("${source[@]:1}")
validpgpkeys=('5FBDF84DD2278DB2B8AD8A5286DED400DDFD7A11') # Isaac Freund <mail@isaacfreund.com>
sha256sums=('af6f536e0c1d049316e17f46941d07b92b035632cfbf38ae6ac6182917d8da13'
            # 'cd7fe3415d4d58685a94fdedd308e9994a37f012828940cfb603461de7f2c6ad'
            # '67fdecba19a5c1b7c2cd98321745fe6cf26a7fdfd5f71140db43c6469b523cc9'
            # '6c960f1805aa187b3217809dab3f9f6e079312c93ab14e3fed5bf948ef4e400f'
            # '1e185423e6b23ed9729614e66751ab7522db4487df4e0dcc7a2b06375aacda23'
            # 'SKIP')
)

# prepare() {
    # zig fetch --global-cache-dir ./zig-global-cache "./${source[1]%%::*}"
    # zig fetch --global-cache-dir ./zig-global-cache "./${source[2]%%::*}"
    # zig fetch --global-cache-dir ./zig-global-cache "./${source[3]%%::*}"
    # zig fetch --global-cache-dir ./zig-global-cache "./${source[4]%%::*}"
    # zig fetch --global-cache-dir ./zig-global-cache "./${source[5]%%::*}"
# }

build() {
  cd $pkgname
  
  DESTDIR="build" zig build \
    -Dllvm --summary all \
    --prefix /usr \
    --search-prefix /usr \
    --build-id=sha1 \
    -Dtarget=native-linux.6.6-gnu.2.40 \
    -Dcpu=baseline \
    -Dpie \
    --build-id=none \
    -Doptimize=ReleaseSafe \
    -Dxwayland

    # --system ../zig-global-cache/p \
    # --global-cache-dir ../zig-global-cache \
}

check() {
  cd $pkgname
  zig build test \
    -Dllvm --summary all \
    --prefix /usr \
    --search-prefix /usr \
    --build-id=sha1 \
    -Dtarget=native-linux.6.6-gnu.2.40 \
    -Dcpu=baseline \
    -Dpie \
    --build-id=none \
    -Doptimize=ReleaseSafe \
    -Dxwayland

    # --global-cache-dir ../zig-global-cache \
    # --system ../zig-global-cache/p \
}

package() {
  cd $pkgname

  cp -a build/* "$pkgdir"

  install -Dm644 README.md -t "$pkgdir/usr/share/doc/$pkgname"
  install -Dm644 contrib/river.desktop -t "$pkgdir/usr/share/wayland-sessions"

  install -d "$pkgdir/usr/share/$pkgname"
  cp -a example "$pkgdir/usr/share/$pkgname"
}
