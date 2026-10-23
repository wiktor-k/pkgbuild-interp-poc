# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=python-lazr.delegates
_name="${pkgname#python-}"
pkgver=2.1.1
pkgrel=3
pkgdesc="Easily write objects that delegate behavior"
arch=(any)
url="https://launchpad.net/lazr.delegates"
license=(LGPL-3.0-only)
depends=(
  python
  python-importlib-metadata
  python-zope-interface
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools
  python-wheel
)
checkdepends=(python-zope-testrunner)
source=(git+https://git.launchpad.net/lazr.delegates#tag=$pkgver
        python-lazr.delegates-2.1.1-remove-pkg_resources.patch)
sha512sums=('6c9aad011b6a7c6d338543a4c710eb82a5907785c4e059c6b1235b52a0ca6437d1b0bda6dc15e3895516bbb0b0cdea0978c5307a140c08e12bfbdcd19e487aea'
            '69c21da13ae27e67043fdbd9aecfefeb12bfee7c64212de86b683a2646b802352bb705152d9e8a98fc2d55e69eb237d06dfad3169c46797ece7212384dbe246c')
b2sums=('580c3bba7805a1a79b8ee0a632d845aeaa35c0744217d75f184e8b80cd7323c38be79fe25b690594f582200d244f56daf8977b4859e377e358f4341708fcd00d'
        '0788bdf3160b067f62826c5a0da868062404ec9846194a51f6baaff61aee1963bfbbd61770a9be023f5b527c548497926875e9456d88349bac1c21dee35376bc')

prepare() {
  cd $_name
  patch -Np1 -i ../python-lazr.delegates-2.1.1-remove-pkg_resources.patch
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
  cd $_name
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CONTRIBUTING,NEWS,README}.rst -t "$pkgdir/usr/share/doc/$pkgname/"
}
