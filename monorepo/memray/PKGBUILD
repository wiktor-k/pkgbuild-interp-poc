# Maintainer: Leonidas Spyropoulos <artafinde@archlinux.org>

pkgname=memray
pkgver=1.20.0
pkgrel=1
pkgdesc="A memory profiler for Python"
arch=('x86_64')
url="https://github.com/bloomberg/memray"
license=('Apache-2.0')
depends=('python' 'libunwind' 'python-rich' 'python-jinja' 'lz4' 'ipython' 'python-textual' 'libelf')
makedepends=('git' 'cmake' 'ninja' 'python-scikit-build-core' 'python-build' 'python-installer' 'python-wheel' 'cython' 'npm' 'python-pkgconfig')
checkdepends=('npm' 'python-pytest' 'python-pytest-textual-snapshot')
source=("git+$url.git#tag=v${pkgver}?signed")
sha512sums=('b36a6eb32d2e6141b2e717109639cae8bdc472acaa5589676fecaa16b1ba24422e08fe973fcc9352821d73e9e0350aff9df1477a9fb7d3471a5e8cf68d6eff8f')
validpgpkeys=('A035C8C19219BA821ECEA86B64E628F8D684696D'   # Pablo Galindo Salgado <pablogsal@gmail.com>
              'F8C887A85001AA63AAE944BC8AB9921EA97C9108')  # Matt Wozniski <godlygeek@gmail.com>

build() {
  cd "${pkgname}"

  export DIST_NAME='Arch Linux'
  export DIST_PACKAGE="$pkgver-$pkgrel"
  export CFLAGS+=" -ffat-lto-objects"
  export CXXFLAGS+=" -ffat-lto-objects"

  make build-js build-vendor
  python -m build --wheel --no-isolation
}

check() {
  cd ${pkgname}
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  # Deselect pytest-textual-snapshot SVG comparison tests which fail due to newer version
  test-env/bin/python -m pytest -vvv \
    --log-cli-level=info -s \
    -k "not TestTUILooks and not test_tui_basic and not test_tui_pause \
        and not test_tui_gradient and not test_merge_threads \
        and not test_unmerge_threads" \
    tests
}

package() {
  cd "${pkgname}"

  python -m installer --destdir="${pkgdir}" dist/*.whl

  install -Dm 644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}

# vim: ts=2 sw=2 et:
