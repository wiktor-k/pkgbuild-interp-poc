# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-configargparse
pkgver=1.8.0
pkgrel=1
pkgdesc='A drop-in replacement for argparse that allows options to also be set via config files and/or environment variables'
arch=('any')
url='https://github.com/bw2/ConfigArgParse'
license=('MIT')
depends=('python')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-setuptools-scm' 'python-wheel')
checkdepends=('python-pytest' 'python-tests' 'python-toml' 'python-yaml')
optdepends=('python-toml: for writing TOML configuration files'
            'python-yaml: for YAML support')
source=("git+https://github.com/bw2/ConfigArgParse.git#tag=v$pkgver")
sha512sums=('2bbc75defa3fc3b7a0cee5854f9ed1d265ac18df5e23f3dec414c0a3eb4c98b3b38c59ead080cd8d82be84607cbb5aa21bfca777d9eb7ac4864f676d0a73852b')

build() {
  cd ConfigArgParse
  python -m build --wheel --no-isolation
}

check() {
  cd ConfigArgParse
  PYTHONPATH=. pytest -v
}

package() {
  cd ConfigArgParse
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
