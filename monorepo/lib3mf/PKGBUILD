# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Jonas Witschel <diabonas@archlinux.org>

pkgbase=lib3mf
pkgname=({,python-}lib3mf)
pkgver=2.5.0
pkgrel=1
pkgdesc='Implementation of the 3D Manufacturing Format file standard'
arch=(x86_64)
url='https://github.com/3MFConsortium/lib3mf'
license=(BSD-2-Clause)
depends=(
  libgcc
  libstdc++
  glibc
  libzip
  util-linux-libs
  zlib
)
makedepends=(
  git
  cmake
  python
  ninja
)
checkdepends=(gtest)
options=(!lto)
source=(
  "$pkgbase::git+$url#tag=v$pkgver"
  fix-install.patch
  __init__.py
)
sha512sums=('3d7cbf0c914521ebacd96599f3a715811388004b4f50de1dba35c9ab9b2fc014a1bd2f1b277cb66bf59627f77239722d099109f513011c5866cf566d0cde2e0a'
            'c4674bc80377126a15422e27925de03a638080f925c55a312472d322ccd529e045b619c239dc37b75915aae7b04bbac857e0a18bd749abd1449d7e30b3101dc6'
            '56cbdd17b95bb98c0d40c294d6a1691594da9630ee0d7ba1e5f0cc4c8be8276edf51af30509b32ba8336a1f24a5b45e3e7436792afa246f3c54d6d8adae47d7c')
b2sums=('11500ef6e1fe79387f34fd643241652da1ac9441504d643343a8774f77629b91c2597865e43a8b536f2e8f99c3154dd17480624fef0c5f146388df5276e2365e'
        '265790be3193ed17261edf049b433daea0091ac6696a57a82de70db7646ace40c755b934e3bd6f7e4806ec17fea25f3d702e9ea4873d088585e9edb242e10c68'
        'cf8f3e695861af2fdacb0d82bce0dc9c9424d8d671ffd37cf59250d3def01f957b57a80b4eb84b7ce0a1ddc339ce0b3a1da5c2bb2e7f2852f9602706b89f9ada')

prepare() {
  cd "$pkgname"

  patch -p1 -i "$srcdir/fix-install.patch"
}

build() {
  local cmake_options=(
    -B build
    -S "$pkgbase"
    -G Ninja
    -D CMAKE_BUILD_TYPE=None
    -D CMAKE_INSTALL_PREFIX=/usr
    -D CMAKE_INSTALL_INCLUDEDIR=include/lib3mf
    -D CMAKE_INSTALL_LIBDIR=lib
    -D CMAKE_POLICY_VERSION_MINIMUM=3.5
    -D USE_INCLUDED_ZLIB=OFF
    -D USE_INCLUDED_LIBZIP=OFF
    -D USE_INCLUDED_SSL=OFF
    -D LIB3MF_TESTS=OFF
    -W no-dev
  )
  cmake "${cmake_options[@]}"
  cmake --build build
}

# NOTE: gtests needs C++14 and with it tests fail
# -D LIB3MF_TESTS=OFF
#check() {
#  ctest --test-dir build --output-on-failure
#}

package_lib3mf() {
  optdepends=('python-lib3mf: Python bindings')
  replaces=(lib3mf-1)

  DESTDIR="$pkgdir" cmake --install build

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" "$pkgbase/LICENSE"
}

package_python-lib3mf() {
  pkgdesc='Python bindings for lib3mf'
  depends=(python lib3mf)

  # provide a python package in site-packages
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  install -vDm644 -t "${pkgdir}${site_packages}/lib3mf" __init__.py
  ln -s /usr/lib/lib3mf.so "${pkgdir}${site_packages}/lib3mf/lib3mf.so"
  ln -s /usr/include/lib3mf/Bindings/Python/Lib3MF.py "${pkgdir}${site_packages}/lib3mf/Lib3MF.py"

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" "$pkgbase/LICENSE"
}
