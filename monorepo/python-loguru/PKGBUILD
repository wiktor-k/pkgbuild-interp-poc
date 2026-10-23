# Maintainer: David Runge <dvzrv@archlinux.org>

_name=loguru
pkgname=python-loguru
pkgver=0.7.3
pkgrel=3
pkgdesc="Python logging made (stupidly) simple"
arch=(any)
url="https://github.com/Delgan/loguru"
license=(MIT)
depends=(python)
makedepends=(
  python-build
  python-flit-core
  python-installer
)
checkdepends=(
  python-colorama
  python-freezegun
  python-pytest
)
source=(
  $_name-$pkgver.tar.gz::$url/archive/refs/tags/$pkgver.tar.gz
  $_name-0.7.3-update_mypy_deps.patch::https://github.com/Delgan/loguru/commit/8bba363a12483b419a27f17212b9368bc3105677.patch?full_index=1
  $_name-0.7.3-disable_mypy_tests.patch::https://github.com/Delgan/loguru/commit/e17479bd0701e8fc0b26981339540599dc224d11.patch?full_index=1
  $_name-0.7.3-fix_exception_modern.patch::https://github.com/Delgan/loguru/commit/84023e2bd8339de95250470f422f096edcb8f7b7.patch?full_index=1
  0001-Bump-flit-core-build-requirement-to-version-4.patch
)
sha256sums=('1cad8860aa0ecf9567125381e4430046526246e075224350a6a624addac05f5e'
            'f6e64ee954877dfae026aa42bb393e14205825632fe949078fa359d5aabecf88'
            '5b36dd7872ab7bd8b62e5fd0fd3501cb4e4cf17364dd9c67f69685ea7ef7243a'
            '1ebc3b7eaf741e542d3a0efca21b30b9845bda50457319abd28bdd224be44f07'
            'd2e4971a35baa36165828630cdcc717d2786b05be1ee8cddd29d8f2e0207040e')
b2sums=('7d7cf167e1350814eea6a358cc00bac217ea6b153ae29ffd70c026f3be63cc126fbc184668ea643ea03416fc8f805bd51502fd8cc9e8d9bcc19099814b8c3fe6'
        '532d0fc1012813274586617fcc2b05dbe12b878f0a155ee5381408b9b6facc02643b5e23813ad13eb232a99a48e3516ab54cf78a21a55a7b89dfe1b8919d3459'
        'ed0705040e89b94414aaaaf25c046c96273d375bf6f6b2816ef2d451472bf074045eb00a06925c8f85de4d9d1db8be058cfff82352b2ace93f9084d31513733e'
        'f946fc2fda917d681754dd9cabc59967b858840d2aa2b19b623f06982dd994826ab30ac1ed78d8012693a9a10b0359edb734930dbf06350390fe53bf44fc135a'
        '855f87b535646939e5dceecbbce9e19cdd6fd36b9f42b2e662a372a9388b6f813e7f666953d53a3ea29f6add7d7558bee13feaf0779cb7fbed6a9d00d90ac7de')

prepare() {
  patch -Np1 -d $_name-$pkgver -i ../$_name-0.7.3-update_mypy_deps.patch
  patch -Np1 -d $_name-$pkgver -i ../$_name-0.7.3-disable_mypy_tests.patch
  patch -Np1 -d $_name-$pkgver -i ../$_name-0.7.3-fix_exception_modern.patch
  patch -Np1 -d $_name-$pkgver -i ../0001-Bump-flit-core-build-requirement-to-version-4.patch
}

build() {
  cd $_name-$pkgver
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    --ignore tests/test_type_hinting.py  # we don't care about type hints
  )

  cd $_name-$pkgver
  pytest "${pytest_options[@]}"
}

package() {
  cd $_name-$pkgver
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 *.rst -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE -t "$pkgdir/usr/share/licenses/$pkgname/"
}
