# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-kaitaistruct
pkgver=0.11
pkgrel=2
pkgdesc="Kaitai Struct declarative parser generator for binary data: runtime library for Python"
arch=('any')
license=('MIT')
url="https://kaitai.io"
depends=('python')
makedepends=('git' 'python-setuptools' 'python-build' 'python-installer')
source=("git+https://github.com/kaitai-io/kaitai_struct_python_runtime.git#tag=v$pkgver")
sha512sums=('472bf20c1ec5122a824b401a464bf6859ef6747a8871771e120179baa3216351cdff40a69cf8c6d2874b5297e022ae58e56f931a8314216c64f7130d719f0cc3')

build() {
  cd kaitai_struct_python_runtime
  python -m build --wheel --no-isolation
}

package() {
  cd kaitai_struct_python_runtime
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
