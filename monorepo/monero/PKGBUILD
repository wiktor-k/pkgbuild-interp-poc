# Maintainer: kpcyrd <kpcyrd[at]archlinux[dot]org>

pkgname=monero
pkgver=0.18.5.4
pkgrel=1
pkgdesc="Monero: the secure, private, untraceable peer-to-peer currency"
license=('BSD-3-Clause')
arch=('x86_64')
url="https://getmonero.org/"
depends=(
  'glibc'
  'libboost_program_options.so'
  'libgcc'
  'libhidapi-libusb.so'
  'libreadline.so'
  'libsodium.so'
  'libssl.so'
  'libstdc++' 'libstdc++.so'
  'libunbound.so'
  'libunwind'
  'libusb'
  'libzmq.so'
  'pcsclite'
  'protobuf'
)
makedepends=(
  'boost'
  'cmake'
  'git'
  'python'
)
source=(
  "${pkgname}::git+https://github.com/monero-project/monero?signed#tag=v${pkgver}"
  "git+https://github.com/miniupnp/miniupnp.git"
  "git+https://github.com/Tencent/rapidjson.git"
  "git+https://github.com/trezor/trezor-common.git"
  "git+https://github.com/tevador/RandomX.git"
  "git+https://github.com/monero-project/supercop.git"
  "fortify-source.patch"
  "monero.sysusers"
  "monero.tmpfiles"
)
sha512sums=('ae6d10759b4bd33eea5b4dd8324166656c0a49d66922d563461632ebc72b733ba84522302b4b6e3553584b1a78fb0215d92b8eb2fbe8b67e2c9434330603d0ee'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP'
            'SKIP'
            'cbb67fee67d053bbc929ad0eeaad6421fae43acf5f1ae4df346333362da7f50adc985a63c27bc4ac2a7c1f226d27162a6cad9908b53866180153a5422038e3db'
            '2ae018bf3660f6eb6ddd1c79ccad1a7545945424d148f2de70607e30a24d39d5ef85f7687980fe0fd0269e029e0c2b8b38c012a78d0f61f01650518632f2d5a5'
            'ae3c300548e6ba5fdf6a203146e5dcaf1955c3e19820c0bb3a73bb70ec5629a3af00d43dac95b6f3832ded5c480e872d829981a63206016d4a64f3940b713357')
validpgpkeys=(
  '8777AB8F778EE89487A2F8E7F4ACA0183641E010' # luigi1111 <luigi1111w@gmail.com>
  'BDA6BD7042B721C467A9759D7455C5E3C0CDCEB9' # Riccardo Spagni <ric@spagni.net>
  '487277A8BD0A209C16B700F3C64552D877C32479' # Alexander Blair (Snipa / Snipa22) <snipa@jagtech.io>
  'E87BD921CDD885C9D78A38C5E45B10DD027D2472' # tobtoht <tob@featherwallet.org>
)

prepare() {
  cd "${pkgname}"
  git submodule init
  git config submodule.external/miniupnp.url "$srcdir/miniupnp"
  git config submodule.external/rapidjson.url "$srcdir/rapidjson"
  git config submodule.external/trezor-common.url "$srcdir/trezor-common"
  git config submodule.external/randomx.url "$srcdir/RandomX"
  git config submodule.external/supercop.url "$srcdir/supercop"
  git -c protocol.file.allow=always submodule update

  patch -Np1 -i ../fortify-source.patch
  sed -e 's|std=c++11|std=c++17|g' -i contrib/depends/protobuf.mk
}

build() {
  cd "${pkgname}"
  CFLAGS+=" -Wno-implicit-function-declaration -Wno-int-conversion" \
  cmake -B build \
    -D ARCH="default" \
    -D CMAKE_BUILD_TYPE=Release \
    -D USE_DEVICE_TREZOR_MANDATORY=ON
  make -C build
}

package() {
  backup=('etc/monerod.conf')

  cd "${pkgname}"
  install -Dm644 "LICENSE" -t "${pkgdir}/usr/share/licenses/${pkgname}"

  install -Dm644 "utils/conf/monerod.conf" "${pkgdir}/etc/monerod.conf"
  install -Dm644 "utils/systemd/monerod.service" "${pkgdir}/usr/lib/systemd/system/monerod.service"
  install -Dm644 "../monero.sysusers" "${pkgdir}/usr/lib/sysusers.d/monero.conf"
  install -Dm644 "../monero.tmpfiles" "${pkgdir}/usr/lib/tmpfiles.d/monero.conf"

  install -Dm755 "build/bin/monero-blockchain-ancestry" \
                 "build/bin/monero-blockchain-depth" \
                 "build/bin/monero-blockchain-export" \
                 "build/bin/monero-blockchain-import" \
                 "build/bin/monero-blockchain-mark-spent-outputs" \
                 "build/bin/monero-blockchain-prune" \
                 "build/bin/monero-blockchain-prune-known-spent-data" \
                 "build/bin/monero-blockchain-stats" \
                 "build/bin/monero-blockchain-usage" \
                 "build/bin/monero-gen-ssl-cert" \
                 "build/bin/monero-gen-trusted-multisig" \
                 "build/bin/monero-wallet-cli" \
                 "build/bin/monero-wallet-rpc" \
                 "build/bin/monerod" \
                 -t "${pkgdir}/usr/bin"
}

# vim: ts=2 sw=2 et:
