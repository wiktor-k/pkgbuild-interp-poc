# Maintainer: Jelle van der Waa <jelle@archlinux.org>

pkgname=python-pika
pkgver=1.4.4
pkgrel=1
pkgdesc="Pure Python RabbitMQ/AMQP 0-9-1 client library"
arch=(any)
url="https://github.com/pika/pika/"
license=('BSD-3-Clause')
depends=('python')
makedepends=('python-build' 'python-installer' 'python-wheel' 'python-setuptools' 'git')
# checkdepends=('python-tornado' 'python-twisted' 'python-gevent')
optdepends=('python-tornado: use with Tornado I/O loop'
            'python-twisted: use with Twisted I/O loop'
            'python-gevent: use with Gevent I/O loop')
source=("git+https://github.com/pika/pika.git#tag=${pkgver}")
sha256sums=('e4d854f85a9e9b9654e53bd12e301a0247369a110519fb46a4a27272974b9236')

build() {
  cd pika
  python -m build --wheel --no-isolation
}

check() {
  cd pika
  # Requires rabbitq node
}

package() {
  cd pika
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -D -m644 LICENSE "${pkgdir}/usr/share/licenses/${pkgname}/LICENSE"
}
