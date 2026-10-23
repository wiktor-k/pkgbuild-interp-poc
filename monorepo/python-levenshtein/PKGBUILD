# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Felix Yan <felixonmars@archlinux.org>
# Contributor: Thomas S Hatch <thatch45@gmail.com>
# Contributor: shamrok <szamrok@gmail.com>
# Contributor: scj <scj archlinux us>

pkgname=python-levenshtein
pkgver=0.27.5
pkgrel=1
pkgdesc='Python extension for computing string edit distances and similarities'
arch=(x86_64)
url='https://github.com/rapidfuzz/Levenshtein'
license=(GPL-2.0-or-later)
depends=(
  glibc
  libgcc
  libstdc++
  python
  python-rapidfuzz
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-scikit-build-core
  cython
  rapidfuzz-cpp
)
checkdepends=(python-pytest)
source=(
  "$pkgname::git+$url#tag=v$pkgver"
  remove-version-constraints.patch
)
sha512sums=('d9654cc2560eb9a380db7397d2b281d4cb83f377d1270a85071c5feb74d6ef323fc8d84a4bfb12a6b06861a92b6ca8b990a83db0a50a882a62d5f37edb3b835d'
            '717cb36ec2e0b66303820dccd771b423cedfd5b30a57b52914cb5f1c2fa84cd064832cacaca5b81350eb83370f0f997e5012fa3c187f0f0b5b16b3d798e5e5b1')
b2sums=('78cd03ac35852094f636ab586cf65df3746f4c03c92df71d6239142bd19075edd009f8a40e9b8dcb1e075203b650a55f8a1d6df5e6b9ebf2afae4fe12c1846c2'
        '27e5d506d31becb2a1150b69e4b1fd683431f1a15ef9964cf09f4a5d4fa955202875a2a34925215b5fa595ec288b90971c01a62f4c4dacbc864b2f36beab3413')

prepare() {
  cd "$pkgname"

  patch -p1 -i "$srcdir/remove-version-constraints.patch"
  sed -e 's|-m cython|-m cython --cplus|g' -i src/Levenshtein/CMakeLists.txt # Fix build with cython 3.2
}
build() {
  cd "$pkgname"

  python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  # install to temporary directory
  python -m installer --destdir="$PWD/tmp_install" dist/*.whl

  PYTHONPATH="$PWD/tmp_install$site_packages" pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl
}
