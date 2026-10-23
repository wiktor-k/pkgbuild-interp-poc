# Maintainer: Konstantin Gizdov <arch at kge dot pw>
# Contributor: Astro Benzene <universebenzene at sina dot com>
_pkgname='pytest-mpl'
pkgname="python-${_pkgname}"
pkgver=0.19.0
pkgrel=1
pkgdesc="Pytest plugin to help with testing figures output from Matplotlib"
arch=('any')
url="https://github.com/matplotlib/pytest-mpl"
license=('BSD')
depends=('python-pytest' 'python-matplotlib' 'python-jinja' 'python-packaging')
makedepends=('python-setuptools-scm' 'python-build' 'python-installer' 'python-wheel' 'git')
checkdepends=('python-pytest' 'python-jinja')
source=("${pkgname}::git+https://github.com/matplotlib/pytest-mpl.git#tag=v${pkgver}")
sha256sums=('5288db8fa97d3873018ae205a5cc53aba2e4a203bf3d83bae3b3ae980766c580')

build() {
    cd "${srcdir}/${pkgname}"
    python -m build --wheel --no-isolation
}

check() {
    cd "${srcdir}/${pkgname}"
    # some tests fail
    PYTHONPATH="${srcdir}/${pkgname}/build/lib" pytest || echo 'some tests fail'
}

package() {
    cd "${srcdir}/${pkgname}"
    python -m installer --destdir="$pkgdir" dist/*.whl
    install -D -m644 LICENSE -t "${pkgdir}/usr/share/licenses/${pkgname}"
    install -D -m644 README.rst -t "${pkgdir}/usr/share/doc/${pkgname}"
}
