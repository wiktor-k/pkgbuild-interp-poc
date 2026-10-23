# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: eolianoe <eolianoe At GoogleMAIL DoT com>
# Contributor: Andrey Mikhaylenko <neithere at gmail dot com>

pkgname=python-pkginfo
_pkgname="${pkgname#python-}"
pkgver=1.13.3
pkgrel=1
pkgdesc='Query metadata from sdists / bdists / installed packages'
arch=(any)
url='https://pypi.python.org/pypi/pkginfo'
license=(MIT)
depends=(python)
makedepends=(
  python-build
  python-installer
  python-wheel
  python-setuptools
)
checkdepends=(
  python-pytest
  python-packaging
  python-wheel
)
source=(
  "$pkgname-$pkgver.tar.gz::https://pypi.io/packages/source/p/pkginfo/pkginfo-$pkgver.tar.gz"
  remove-pkg_resources.patch
)
sha512sums=('38390a841a8576c2b2a68eb9ed87d9d54d9c696f7f3dfe8d5d91d6e751d4655cdc6dc7f353b2d5eb72abffd00d4f0b17a32773720b651556feb08218d7c4942d'
            'b568f86ed50c02f6f6c08b5f79327d9729278a4677d117a2b1f255951c281e0348d24eeb23bde87b9a40e60d684620bec95c404f013e6756c8cee89aec575f8f')
b2sums=('0f5e14c9715b7ccbc22c6a3ce189b25db511b8498513514c4f5dfd108a15a68d12e2a30376845a66d5a3168d98b3d743b761e65a76420e5466bef57b1b17f7c8'
        'e7ef432acb7db3f16dc9f6b011a713dfe4ac281ff5a895d18d7b287360cd4281e18650a0b3a153e7aa58a3d05c874fe8a8e26b6b12a83df012d0e915c675dbf3')

prepare() {
  cd "$_pkgname-$pkgver"

  patch -p1 -i "$srcdir/remove-pkg_resources.patch"
}

build() {
  cd "$_pkgname-$pkgver"

  python -m build --wheel --no-isolation
}

check() {
  cd "$_pkgname-$pkgver"

  local pytest_options=(
    -vv
    # HACK: needs upstream big report
    -k 'not test_installed_ctor_w_dist_info'
    --deselect tests/checkers.py
    # ?
    --deselect pkginfo/tests/test_installed.py::test_installed_ctor_w_package
    --deselect pkginfo/tests/test_installed.py::test_installed_ctor_w_package_and_metadata_version
    --deselect pkginfo/tests/test_installed.py::test_installed_ctor_w_name
    --deselect pkginfo/tests/test_installed.py::test_installed_ctor_w_name_and_metadata_version
    --deselect pkginfo/tests/test_utils.py::test_get_metadata_w_module
    --deselect pkginfo/tests/test_utils.py::test_get_metadata_w_module_and_metadata_version
    --deselect pkginfo/tests/test_utils.py::test_get_metadata_w_package_name
    --deselect pkginfo/tests/test_utils.py::test_get_metadata_w_package_name_and_metadata_version
  )
  pytest "${pytest_options[@]}"
}

package() {
  cd "$_pkgname-$pkgver"

  python -m installer --destdir="$pkgdir" dist/*.whl

  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE.txt
}
