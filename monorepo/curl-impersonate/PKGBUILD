# Maintainer: George Rawlinson <grawlinson@archlinux.org>

pkgname=curl-impersonate
pkgver=2.2.2
pkgrel=1
pkgdesc='A build of curl that impersonates four major browsers: Chrome, Edge, Safari & Firefox'
arch=(x86_64)
url='https://github.com/lexiforest/curl-impersonate'
license=(MIT)
depends=(
  glibc
  libgcc libgcc_s.so
  libstdc++ libstdc++.so
  bash
)
makedepends=(
  git
  cmake
  go
  ninja
  unzip
  chrpath
)
_brotli_version=1.2.0
_boringssl_commit=156c7b75ae9b8c3b3f847acf264f17594c3859fb
_nghttp2_version=1.63.0
_ngtcp2_version=1.20.0
_nghttp3_version=1.15.0
_curl_version=curl-8_21_0
_zlib_version=1.3.1
_zstd_version=1.5.7
_libidn2_version=2.3.7
source=(
  "$pkgname::git+$url#tag=v$pkgver"
  "brotli-${_brotli_version}.tar.gz::https://github.com/google/brotli/archive/refs/tags/v${_brotli_version}.tar.gz"
  "boringssl-${_boringssl_commit}.zip::https://github.com/google/boringssl/archive/${_boringssl_commit}.zip"
  "nghttp2-${_nghttp2_version}.tar.bz2::https://github.com/nghttp2/nghttp2/releases/download/v${_nghttp2_version}/nghttp2-${_nghttp2_version}.tar.bz2"
  "ngtcp2-${_ngtcp2_version}.tar.bz2::https://github.com/ngtcp2/ngtcp2/releases/download/v${_ngtcp2_version}/ngtcp2-${_ngtcp2_version}.tar.bz2"
  "nghttp3-${_nghttp3_version}.tar.bz2::https://github.com/ngtcp2/nghttp3/releases/download/v${_nghttp3_version}/nghttp3-${_nghttp3_version}.tar.bz2"
  "${_curl_version}.tar.gz::https://github.com/curl/curl/archive/${_curl_version}.tar.gz"
  "zlib-${_zlib_version}.tar.gz::https://github.com/madler/zlib/releases/download/v${_zlib_version}/zlib-${_zlib_version}.tar.gz"
  "zstd-${_zstd_version}.tar.gz::https://github.com/facebook/zstd/releases/download/v${_zstd_version}/zstd-${_zstd_version}.tar.gz"
  "libidn2-${_libidn2_version}.tar.gz::https://ftp.gnu.org/gnu/libidn/libidn2-${_libidn2_version}.tar.gz"
  no-download.patch
)
noextract=(
  "brotli-${_brotli_version}.tar.gz"
  "boringssl-${_boringssl_commit}.zip"
  "nghttp2-${_nghttp2_version}.tar.bz2"
  "ngtcp2-${_ngtcp2_version}.tar.bz2"
  "nghttp3-${_nghttp3_version}.tar.bz2"
  "${_curl_version}.tar.gz"
  "zlib-${_zlib_version}.tar.gz"
  "zstd-${_zstd_version}.tar.gz"
  "libidn2-${_libidn2_version}.tar.gz"
)
sha512sums=('c59e564a07357f6f2b67eb961f5521c16bf7aa3e47e6813265b6707984fe821793f199511855f1afa2f1224aea628a7d6315366c1c1d840487959bf74ddf672d'
            'f94542afd2ecd96cc41fd21a805a3da314281ae558c10650f3e6d9ca732b8425bba8fde312823f0a564c7de3993bdaab5b43378edab65ebb798cefb6fd702256'
            '1bd18f7344d6b30ff58f7e4319d7932ba915f235ee7552a997fc59977a863fc66ab3bbcc5020584cf0ef0b17af0bc2bcc87e083c7cd71993c9fc78e5af469767'
            'a328b4642c6ca4395adfcaaf4e6eb6dbd39fa7bd86f872a76260af59a5a830e0ff5ad015865d6bc00e0baa8e4d0d9a67b4b97e9d78e5e05d1c53522364e5e235'
            'fead05300f91fe79bbacb1f57ca7f5e390ec0af5f847bcf2800434e0fa9a895851b71cc9774b0ac04cdd4d660bb1c02ddfa88cb08287649f57ce327a434b4262'
            'e2de0402f3b75f081f421052e8fc72d7a679fbab67b8b1ba383bbfc61c90e7ff58e475c563747712a680e367af237c47f3f890326a55440bfdd2e51d0ea1fb6c'
            '0ab6c99c3d5b86fb65c526db517c3159b11db2f8d82552d635c4887059c0602288603c93b754ce0ec543ea2f275122ccec2c8dcd866c2611b5b949c728ee72df'
            '580677aad97093829090d4b605ac81c50327e74a6c2de0b85dd2e8525553f3ddde17556ea46f8f007f89e435493c9a20bc997d1ef1c1c2c23274528e3c46b94f'
            'b4de208f179b68d4c6454139ca60d66ed3ef3893a560d6159a056640f83d3ee67cdf6ffb88971cdba35449dba4b597eaa8b4ae908127ef7fd58c89f40bf9a705'
            'eab5702bc0baed45492f8dde43a4d2ea3560ad80645e5f9e0cfa8d3b57bccd7fd782d04638e000ba07924a5d9f85e760095b55189188c4017b94705bef9b4a66'
            'bc4cf3aa0a1be6bde148489756d495a92b3d079494d214b0a7905bbd5fecd58710f507cd43685be6900a1d622d4ad34825d73494496b05c6ca024af36f414406')
b2sums=('69cbe84a8cd95036777458639cbec42d9613ef10d83ed53c21c8fc1d5dfdacc8c3920efd8cddae56f1ba96908ab0356a9104bc1276260bc5a668f9fd122c194a'
        'e38ea35f4abe47b2113ccc0778bb7b0be77dcf4d26ce248347ecc4cfdea7ba40d17e0b175cad6f82055d07b1ec9976ffea649f59065866db2c36116c77e59824'
        '65a466046a3e24e36d810afc18c700e6e41bbf1f683be858cf973e85a48126ec1562cfd9f7f91dba50f4c78591bdfc0af823587c30d75bc79d9bd2991a3d2999'
        '8caab3602d79abe03f0b266feba07edbdcc95288824d41b39c672a74a9ef9c97e9421138eba4fc04c87aaaa0f26a7ce3779831825629f421f293481fa9f315e3'
        'cb5fa78f07c8fe1dfe5890894e1f7811ae5750f7afc77bd6ff83b7086cd0654882f6acc58f2611af2a69827125a6f58d6e11721958a1d2ebd7b4fd092ba6a5eb'
        '95064a8c9025822937af77da74b6f9b3bb6bef012678d20a2daf20c16b9f242c0a637f2dac009da82addb359200a72614efa1e2cbb2b464ae164d58c4d6f3f02'
        '3e4df97e213a5aa56ef144cdfd085ea51b2902cf09d6006315e8036156eded548814a43983cd2eea9a653d635219675a8223682f2c98ce4379a8d6f5cd14f63c'
        '872ef9cc41ba57e9809ab5714ffa15cc9d3ef6c4a948c57107f800fc373bcfe2475136407203cb9aa33b189994336f36f2757b7a582e065ad477bfd49260184b'
        'ef6bcf13bbb79edce2de21ef4860b01dcb4d2dd47ebd03657d5331c3bb1a207ef128e99ed907a32e787dca7bb51df225ebd8abb2097d03d02c003ca713e1e271'
        '40fde88e6be1e4665738855608a65c949834728463a3fb466192403f1076e99f7d399b9631308ff6776a28915251a0136ddfe6a67e8e010c7ccdb7fc64125d39'
        '074ef882f97baec253c7aee9db9af483018f9331a8734fa9088081ee88fe983b7b9fe18eabf1b451ceddb842834eb12ca0a0c2385a872fc44752a9a324a5babe')

prepare() {
  cd "$pkgname"

  # backport fix for cmake files being installed
  # https://github.com/lexiforest/curl-impersonate/issues/289
  git cherry-pick --no-commit a0bdfdd8ec4a8c87013a0fd014746d0447ee2910

  mkdir -p build/deps/downloads
  mv "$srcdir"/*.{tar*,zip} build/deps/downloads

  # do not download dependencies
  patch -p1 -i "$srcdir/no-download.patch"
}

build() {
  cd "$pkgname"

  make prepare-libidn2

  local cmake_options=(
    -B build
    -S .
    -G Ninja
    -W no-author
    -D CMAKE_BUILD_TYPE=None
    -D CMAKE_INSTALL_PREFIX=/usr
  )
  
  cmake "${cmake_options[@]}"
  
  cmake --build build
}

package() {
  cd "$pkgname"

  DESTDIR="$pkgdir" cmake --install build

  # remove unnecessary files
  rm -rf "$pkgdir/usr/include"

  # remove unsafe rpath
  chrpath --delete \
    "$pkgdir/usr/lib/libcurl-impersonate.so" \
    "$pkgdir/usr/bin/curl-impersonate"

  # license
  install -vd "$pkgdir/usr/share/licenses/$pkgname"
  mv "$pkgdir"/usr/LICENSE* "$pkgdir/usr/share/licenses/$pkgname"
}
