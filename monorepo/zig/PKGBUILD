# Maintainer: Daurnimator <daurnimator@archlinux.org>
# Contributor: Marc Tiehuis <marctiehuis@gmail.com>

pkgname=zig
pkgver=0.17.0
pkgrel=1
pkgdesc='a general-purpose programming language and toolchain for maintaining robust, optimal, and reusable software'
arch=('x86_64')
url='https://ziglang.org/'
license=('MIT')
options=('!lto')
depends=('clang22' 'lld22' 'llvm22-libs')
makedepends=('cmake' 'llvm22')
checkdepends_x86_64=('lib32-glibc')
source=("https://ziglang.org/download/$pkgver/zig-$pkgver.tar.xz"
        "skip-futex2-test.patch")
sha256sums=('b6c7f1728f043700d6529bac980800792f824256a9d2f1839b3d62beed0b8abd'
            'eb30e0eb00e6ced4c99383f0658a0351f42882e303300ed1828d162d27171cd0')

prepare() {
    cd "$pkgname-$pkgver"

    patch -p1 -i ../skip-futex2-test.patch
}

build() {
    cd "$pkgname-$pkgver"

    local cmake_vars=(
        CMAKE_INSTALL_PREFIX=/usr

        # The zig CMakeLists uses build type Debug if not set
        # override it back to None so makepkg env vars are respected
        CMAKE_BUILD_TYPE=None

        ZIG_PIE=ON
        ZIG_SHARED_LLVM=ON
        ZIG_USE_LLVM_CONFIG=ON

        ZIG_TARGET_TRIPLE=native-linux.6.6-gnu.2.40
        ZIG_TARGET_MCPU=baseline
    )
    cmake -B build "${cmake_vars[@]/#/-D}" .
    cmake --build build
}

check() {
    cd "$pkgname-$pkgver"
    # ugly workaround until test target is provided
    # https://github.com/ziglang/zig/issues/14240
    DESTDIR="./testinstall" cmake --install build
    ./testinstall/usr/bin/zig build test \
      -Dconfig_h=build/config.h \
      -Dversion-string="$pkgver" \
      -Dstatic-llvm=false \
      -Denable-llvm=true \
      -Dskip-non-native=true
}

package() {
    cd "$pkgname-$pkgver"

    install -Dm644 LICENSE "$pkgdir/usr/share/licenses/$pkgname/LICENSE"

    DESTDIR="$pkgdir" cmake --install build
}
