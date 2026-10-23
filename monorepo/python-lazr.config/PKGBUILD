# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-lazr.config
_name="${pkgname#python-}"
pkgver=3.1
pkgrel=4
pkgdesc="Create configuration schemas, and process and validate configurations"
arch=(any)
url="https://launchpad.net/lazr.config"
license=(LGPL-3.0-only)
depends=(
  python
  python-lazr.delegates
  python-zope-interface
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(
  python-zope-testrunner
)
provides=(python-lazr-config)
replaces=(python-lazr-config)
source=(git+https://git.launchpad.net/lazr.config#tag=$pkgver
        python-lazr.config-3.1-remove-pkg_resources.patch)
sha512sums=('ee5e10ac0be62bedfd69d01dfa364b0eb75e6ad7d291d096d13d15b902cac33310d5a8c1812656b04edcb3f5489e164fa0cbac1cf42aa00ab85b69762ccc85eb'
            '5479523849f9193f0fd481f04eeb989da1c20af7e1ea0b00926696e1435257a9bac6e17f825c281b89f864721cd5c43d6fd6558126c882c01f971899bef81fee')
b2sums=('2b49fadf20e0066e2a25a783b60f25e59730e00411c1f25f3d7596a0061664b44b90fb6dd8aeaaf77809fece70ee7f34f2207d3255c0c521b5be60f0c44a685a'
        '9f4cb8ba550dbe74124471f396b542a5e25d815a4d5e7439340864a7071c79279905e33379ba5588fbb5fd1882b79e3919398978ba7e74fefdeede689f83971e')

prepare() {
  cd $_name
  patch -Np1 -i ../python-lazr.config-3.1-remove-pkg_resources.patch
}

build() {
  cd $_name
  python -m build --wheel --no-isolation
}

check() {
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name
  python -m installer --destdir=test_dir dist/*.whl
  export PYTHONPATH="$PWD/test_dir/$site_packages:$PYTHONPATH"
  python -m zope.testrunner --test-path src --tests-pattern ^tests
}

package() {
  local site_packages=$(python -c "import site; print(site.getsitepackages()[0])")

  cd $_name
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {NEWS,README}.rst -t "$pkgdir/usr/share/doc/$pkgname/"
  rm -frv "$pkgdir/$site_packages/lazr/config/tests"
}
