# Maintainer: Daurnimator <daurnimator@archlinux.org>
# Contributor: Andrea Feletto <andrea@andreafeletto.com>

pkgname=river
pkgver=0.5.0
pkgrel=1
pkgdesc='a non-monolithic Wayland compositor'
arch=('x86_64')
url="https://isaacfreund.com/software/river/"
license=('0BSD' 'CC-BY-SA-4.0' 'GPL-3.0-only' 'MIT')
depends=(
    'glibc'
    'libevdev'
    'libinput'
    'libxkbcommon'
    # 'mesa'
    'pixman'
    # 'sh'
    'wayland'
    'wlroots0.20'
    'xorg-xwayland'
)
makedepends=(
    'git'
    'scdoc'
    'wayland-protocols'
    'zig'
    # https://codeberg.org/ziglang/zig/issues/31272
    'clang'
)
# optdepends=('polkit: access seat through systemd-logind')
# provides=('river-window-management-v1')

source=(git+https://codeberg.org/river/river.git#tag=v${pkgver}?signed)
#
# disable source caching for now: https://codeberg.org/ziglang/zig/issues/31866
# 
# PACKAGING.md -> build.zig.zon
# source+=(zig-pixman-0.3.0.tar.gz::https://codeberg.org/ifreund/zig-pixman/archive/v0.3.0.tar.gz
         # zig-wayland-0.6.0.tar.gz::https://codeberg.org/ifreund/zig-wayland/archive/v0.6.0.tar.gz
         # zig-wlroots-0.20.1.tar.gz::https://codeberg.org/ifreund/zig-wlroots/archive/v0.20.1.tar.gz
         # zig-xkbcommon-0.4.0.tar.gz::https://codeberg.org/ifreund/zig-xkbcommon/archive/v0.4.0.tar.gz
         # translate_c::git+https://codeberg.org/ziglang/translate-c.git#commit=57c559cf581b1fcad90494eda219f98abeb155ce)
noextract=("${source[@]:1}")
conflicts=('river-classic')
validpgpkeys=('5FBDF84DD2278DB2B8AD8A5286DED400DDFD7A11') # Isaac Freund <mail@isaacfreund.com>
sha256sums=('4ebf58cb64a9cb21838c241233edcde0d2274c2f8175d222cb147b993b2d7e7d')

# prepare() {
#     zig fetch --global-cache-dir ./zig-global-cache "./${source[1]%%::*}"
#     zig fetch --global-cache-dir ./zig-global-cache "./${source[2]%%::*}"
#     zig fetch --global-cache-dir ./zig-global-cache "./${source[3]%%::*}"
#     zig fetch --global-cache-dir ./zig-global-cache "./${source[4]%%::*}"
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
    # --global-cache-dir ../zig-global-cache \
    # --system ../zig-global-cache/p \
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

  # add license + exception
  install -m644 -Dt "${pkgdir}/usr/share/licenses/${pkgname}" LICENSES/{0BSD.txt,MIT.txt}
}
