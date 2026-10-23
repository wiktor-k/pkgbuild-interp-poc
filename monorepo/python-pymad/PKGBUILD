# Maintainer: George Rawlinson <grawlinson@archlinux.org>
# Contributor: Alexander F. Rødseth <xyproto@archlinux.org>
# Contributor: Mark Rosenstand <mark@archlinux.org
# Contributor: Corrado Primier <cp663137@stud.polimi.it>

pkgname=python-pymad
pkgver=0.11.3
pkgrel=4
pkgdesc='Python wrapper for the MPEG Audio Decoder (MAD) library'
arch=(x86_64)
url='https://github.com/jaqx0r/pymad'
license=(GPL-2.0-only)
depends=(
  glibc
  libmad
  python
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-setuptools-scm
  python-wheel
)
checkdepends=(python-pytest)
source=(
  "$pkgname::git+$url#tag=v$pkgver"
  pyproject-license.patch
)
sha512sums=('e7d39cbfd2dcfce7f7aa075d81e99ceb157a2f6a7b25d4ce54d8e639a7cd3d5ac4a4091e7699feb687b67255beb9d8f47513e0a2396e47dda67312656c8127b0'
            '6904874bb3d0947004976837f88dd0c71c1b17376a159448e318ead27c6f86cb260e1e37f3acb1ee9ffc39bf79de09a78a01f49d0671d613ad832dbef34f3cef')
b2sums=('c8d9b4a265bacc7c61cbd2a4125e697f76292627d48b00a3fb3aa6b208ea0ee2f8a95cb0ed8bb8e33d8e85c12bf1a770f6ed565b175c7d34d67de3edf8afc0a3'
        '8b2fd5ec2fbfdab6d89444c512907f11870faa57db9046015784f7b8dc04f86e6bd04ee439a934f839df2a9af2b26cdae2978d468153934b8ad6000924c58213')

prepare() {
  cd "$pkgname"

  # avoid usage of deprecated license format
  patch -p1 -i "$srcdir/pyproject-license.patch"
}

build() {
  cd "$pkgname"

  SETUPTOOLS_SCM_PRETEND_VERSION="$pkgver" python -m build --wheel --no-isolation
}

check() {
  cd "$pkgname"

  # temporary install
  python -m installer --destdir="$(pwd)/tmp" dist/*.whl
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")
  export PYTHONPATH="$(pwd)/tmp/$site_packages"

  pytest -v
}

package() {
  cd "$pkgname"

  python -m installer --destdir="$pkgdir" dist/*.whl
}
