# Maintainer: Alexander Epaneshnikov <alex19ep@archlinux.org>

pkgname=python-immutabledict
pkgver=4.3.1
pkgrel=1
pkgdesc="Immutable wrapper around dictionaries (a fork of frozendict)"
arch=('any')
url=https://github.com/corenting/immutabledict
license=('MIT')
depends=('python')
checkdepends=('python-pytest')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-poetry-core')
source=("$pkgname-$pkgver.tar.gz::https://github.com/corenting/immutabledict/archive/refs/tags/v$pkgver.tar.gz")
sha512sums=('c9d1b647a3e56f9675b0029605b319196b04ac18bd02bfbc544e0e9a23444a7fe14ca986c378ba919bda6dc9d78ac6447c39ab60e7a185b3415832861ca4d1dd')

build() {
    cd "immutabledict-$pkgver"
    python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
    cd "immutabledict-$pkgver"
    python -m venv --system-site-packages test-env
    test-env/bin/python -m installer dist/*.whl
    PATH="$PWD/test-env/bin:$PATH" test-env/bin/python -m pytest
}

package() {
    cd "immutabledict-$pkgver"
    python -m installer --destdir="$pkgdir" dist/*.whl
    install -Dm0644 -t "$pkgdir/usr/share/licenses/$pkgname/" LICENSE
}
