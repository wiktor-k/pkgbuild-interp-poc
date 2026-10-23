# Maintainer: David Runge <dvzrv@archlinux.org>

_name=django-mailman3
pkgname=python-django-mailman3
pkgver=1.3.15
pkgrel=5
pkgdesc="Libraries and templates to Django-based interfaces interacting with Mailman"
arch=(any)
url="https://gitlab.com/mailman/django-mailman3"
license=(GPL-3.0-or-later)
depends=(
  python
  python-cryptography  # for python-django-allauth social account support
  python-django
  python-django-allauth
  python-django-gravatar
  python-mailmanclient
  python-openid  # for python-django-allauth openid support
  python-pyjwt  # for python-django-allauth social account support
  python-requests  # for python-django-allauth social account support
  python-requests-oauthlib  # for python-django-allauth social account support
)
makedepends=(
  git
  python-build
  python-installer
  python-pdm-backend
  python-wheel
)
checkdepends=(python-pytest)
source=(
  "git+$url?signed#tag=v$pkgver"
  "hack-incompatible-allauth-test.patch"
  "bump-django-version.patch"
)
sha512sums=('cfe2cbdc254ba089962969f6f57e4c8761ea60ed491930789260e24aff22d864cd41b774c13924df29d295a8ab5d4f929a4c23928cd927ad40054d376882794c'
            '87d102617a0d8bae603c925074b132abdb4cdd75b062a75776242ea0e49aba878440492a4b8d755a445e648e36c2ebedb79cd4d8ecbc98fa66bd3fa5cd38b4c2'
            '33cb9db08bc8eff0110909fd161da4cdcfeb74d1cf1f565043877bc483a5bb22b903b4e00de95df5843e1721fc6ab77999c82a73b6baa8075bcf5c0b62be6c2a')
b2sums=('320812e7d597fa5616305a6d868a75f22819ca2050defa0572eea5ec687cd862628f88a97d27270e594de270e4c385ff69bdcaf00634d455d22c4f1315eee6e5'
        '37dc4c1371b61aa95de997409554a91cb2e31b0a6928457c1c62cdd1ef0af65438908faa86de66f07d4c35851101bf96cfa82c8b549ab8407ec668b1f7507d62'
        '8a0600c76e46bd74232fc099c2ab4a51e1ee3fbcece8393120c4d20d49c0507e75221de6d53cbe68cdc5e5f53900a5572ac299657cfaf857b8f9f38acbc0c8bb')
validpgpkeys=('541EA0448453394FF77A0ECC9D9B2BA061D0A67C') # Abhilash Raj <raj.abhilash1@gmail.com>

pkgver() {
  cd $_name
  git describe | sed 's/\([^-]*-g\)/r\1/;s/v//g;s/-/./g'
}

prepare() {
  cd $_name/django_mailman3/
  # https://gitlab.com/mailman/django-mailman3/-/issues/82
  patch -Np1 -i ${srcdir}/hack-incompatible-allauth-test.patch

  # Bump django version to 5.1
  cd $srcdir/$_name
  patch -Np0 -i ${srcdir}/bump-django-version.patch
}

build() {
  cd $_name
  python -m build --wheel --no-isolation
}

check() {
  cd $_name
  export PYTHONPATH="build:$PYTHONPATH"
  django-admin test --settings=django_mailman3.tests.settings_test django_mailman3
}

package() {
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 README.rst -t "$pkgdir/usr/share/doc/$pkgname"
}
