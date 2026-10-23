# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>

pkgname=monero-gui
pkgver=0.18.5.4
pkgrel=1
pkgdesc="Qt GUI wallet for Monero: the secure, private, untraceable peer-to-peer currency"
license=('BSD-3-Clause')
arch=('x86_64')
url="https://getmonero.org/"
depends=(
  'glibc'
  'hicolor-icon-theme'
  'libboost_program_options.so'
  'libgcc'
  'libgcrypt'
  'libhidapi-libusb.so'
  'libsodium.so'
  'libssl.so'
  'libstdc++' 'libstdc++.so'
  'libunbound.so'
  'libunwind'
  'libusb'
  'libx11'
  'monero'
  'pcsclite'
  'qt5-base'
  'qt5-declarative'
  'qt5-graphicaleffects'
  'qt5-quickcontrols'
  'qt5-quickcontrols2'
  'qt5-svg'
  'qt5-xmlpatterns'
)
makedepends=(
  'boost'
  'cmake'
  'git'
  'python'
  'qt5-tools'
)
source=(
  "git+https://github.com/monero-project/monero-gui?signed#tag=v${pkgver}"
  "git+https://github.com/monero-project/monero"
  "git+https://github.com/dlbeer/quirc.git"
  "git+https://github.com/miniupnp/miniupnp.git"
  "git+https://github.com/Tencent/rapidjson.git"
  "git+https://github.com/trezor/trezor-common.git"
  "git+https://github.com/tevador/RandomX.git"
  "git+https://github.com/monero-project/supercop.git"
)
sha512sums=('936895f6276f47a1ab669dd96f018a178430776228badf83d62d1f6f73c7eda746b796970216462c65f212007f149420fc0444a9eba06f3bb403f0736cadfa3b'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP')
validpgpkeys=(
  '8777AB8F778EE89487A2F8E7F4ACA0183641E010' # luigi1111 <luigi1111w@gmail.com>
  '487277A8BD0A209C16B700F3C64552D877C32479' # Alexander Blair (Snipa / Snipa22) <snipa@jagtech.io>
  'E87BD921CDD885C9D78A38C5E45B10DD027D2472' # tobtoht <tob@featherwallet.org>
)

prepare() {
  cd "${pkgname}"
  git submodule init
  git config submodule.monero.url "$srcdir/monero"
  git config submodule.external/quirc.url "$srcdir/quirc"
  git -c protocol.file.allow=always submodule update

  sed -i '/<releases>/,/<\/releases>/d' share/org.getmonero.Monero.metainfo.xml

  cd monero
  git config submodule.external/miniupnp.url "$srcdir/miniupnp"
  git config submodule.external/rapidjson.url "$srcdir/rapidjson"
  git config submodule.external/trezor-common.url "$srcdir/trezor-common"
  git config submodule.external/randomx.url "$srcdir/RandomX"
  git config submodule.external/supercop.url "$srcdir/supercop"
  git -c protocol.file.allow=always submodule update
}

build() {
  cd "${pkgname}"
  mkdir -p build && cd build
  CFLAGS+=" -Wno-implicit-function-declaration -Wno-int-conversion" \
  cmake \
    -D CMAKE_BUILD_TYPE=Release \
    -D ARCH=default \
    -D WITH_DESKTOP_ENTRY=OFF \
    -D WITH_UPDATER=OFF \
    -D USE_DEVICE_TREZOR_MANDATORY=ON \
    ../
  make
}

package() {
  cd "${pkgname}"
  install -Dm755 build/bin/monero-wallet-gui -t "${pkgdir}/usr/bin"
  install -Dm644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"

  install -Dm644 share/org.getmonero.Monero.desktop -t "${pkgdir}/usr/share/applications"
  install -Dm644 share/org.getmonero.Monero.metainfo.xml -t "${pkgdir}/usr/share/metainfo"
  for x in 16 24 32 48 64 96 128 256; do
    install -Dm644 "images/appicons/${x}x${x}.png" "${pkgdir}/usr/share/icons/hicolor/${x}x${x}/apps/org.getmonero.Monero.png"
  done
}

# vim: ts=2 sw=2 et:
