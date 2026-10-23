# Maintainer: David Runge <dvzrv@archlinux.org>

pkgname=rst2pdf
pkgver=0.105
pkgrel=1
pkgdesc="Use a text editor. Make a PDF"
arch=(any)
url="https://github.com/rst2pdf/rst2pdf"
license=(MIT)
depends=(
  python
  python-docutils
  python-jinja
  python-packaging
  python-pillow
  python-pygments
  python-reportlab
  python-yaml
)
makedepends=(
  git
  python-build
  python-installer
  python-setuptools-scm
  python-wheel
)
# TODO: package python-xhtml2pdf
# https://github.com/xhtml2pdf/xhtml2pdf/issues/560
checkdepends=(
  graphviz
  plantuml
  python-aafigure
  python-matplotlib
  python-pymupdf
  python-pytest
  python-pytest-xdist
  python-sphinx
  python-svglib
)
# TODO: package python-xhtml2pdf
optdepends=(
  'plantuml: for UML diagram support'
  'python-aafigure: for ASCII image support'
  'python-pyphen: for hyphenation support'
  'python-matplotlib: for math support'
  'python-sphinx: for sphinx support'
  'python-svglib: for SVG support'
)
source=(
  "git+$url.git#tag=$pkgver"
)
sha512sums=('299cc9b7f164aeb1b2bbf20efb67103f9e60d06fbe6aaea3a0bb9ddfb9ee7b01f6322fd3bc3e8264f0fdf90fd666338f997fe491989c0f68a4c835a9bc191a7e')
b2sums=('be3416aea5e8955c13a2f44bf7ee3cad3add2d6bfb94a677f5329ecda2682f29af3f9b9656e2969626de16274d054de41a0e54979804db954b7de56805485a39')

build() {
  cd $pkgname
  python -m build --wheel --no-isolation
}

check() {
  local pytest_options=(
    -vv
    --numprocesses=auto
    # https://github.com/rst2pdf/rst2pdf/issues/1067
    --deselect=tests/input/test_raw_html.rst::test_raw_html
  )

  cd $pkgname
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  export PYTHONPATH="$(test-env/bin/python -c "import site; print(site.getsitepackages()[0])")"
  export PATH="$PATH:$PWD/test-env/bin"
  test-env/bin/python -Pm pytest "${pytest_options[@]}"
}

package() {
  cd $pkgname
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -vDm 644 {CHANGES,MAINTAINERS,README}.rst -t "$pkgdir/usr/share/doc/$pkgname/"
  install -vDm 644 LICENSE.txt -t "$pkgdir/usr/share/licenses/$pkgname/"
}
