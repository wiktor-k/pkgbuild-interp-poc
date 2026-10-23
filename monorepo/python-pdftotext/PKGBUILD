# Maintainer: Daniel M. Capella <polyzen@archlinux.org>
# Contributor: Shane Stone <shanewstone at gmail>

pkgname=python-pdftotext
_name=${pkgname#python-}
pkgver=4.0.0
pkgrel=1
pkgdesc='Simple PDF text extraction'
arch=('x86_64')
url='https://sr.ht/~jalan/pdftotext/'
license=('MIT')
depends=('poppler' 'python')
makedepends=(
  'git'
  'python-build'
  'python-installer'
  'python-setuptools'
  'python-wheel'
)
source=("git+https://git.sr.ht/~jalan/pdftotext#tag=v$pkgver")
b2sums=('36030371b4d94a399a06e043a9d45315672b54bb0b89a44c06a508b236be7c10ba32f5c6b3aef66689b6f8f65b15697bba7e31c9d495c3aacccbbbaed9203956')

build() {
  cd "$_name"
  python -m build --wheel --skip-dependency-check --no-isolation
}

check() {
  cd "$_name"
  python -m venv --system-site-packages test-env
  test-env/bin/python -m installer dist/*.whl
  test-env/bin/python -m unittest discover
}

package() {
  depends+=(
    glibc
    libgcc libgcc_s.so
    libstdc++ libstdc++.so
  )

  cd "$_name"
  python -m installer --destdir="$pkgdir" dist/*.whl

  # license
  install -vDm644 -t "$pkgdir/usr/share/licenses/$pkgname" LICENSE
}
