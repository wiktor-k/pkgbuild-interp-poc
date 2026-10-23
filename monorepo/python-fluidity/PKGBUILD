# Maintainer: Felix Yan <felixonmars@archlinux.org>

pkgname=python-fluidity
pkgver=0.2.1
_commit=468627ea13ed4c3043dc67ccc1ff16089fdb7b25
pkgrel=8
pkgdesc="Fluidity: state machine implementation for Python objects"
url="https://github.com/nsi-iff/fluidity"
license=('MIT')
arch=('any')
depends=('python')
makedepends=('git' 'python-build' 'python-installer' 'python-setuptools' 'python-wheel')
checkdepends=('python-should-dsl')
source=("git+https://github.com/nsi-iff/fluidity.git#commit=$_commit")
sha512sums=('2ff7c490c0503fd9019afe176e97524b2644d5d6b05be4a0dc49c3455212e4bf3f1bd4caa6206b4bf068fedabe7f83514d3baedcb0189449dee77074315792ef')

build() {
  cd fluidity
  python -m build --wheel --no-isolation
}

check() {
  cd fluidity
  (
    export PYTHONPATH="$PWD"
    # Copied from https://github.com/nsi-iff/fluidity/blob/master/tox.ini
    python spec/callable_support_spec.py
    python spec/callback_order_spec.py
    python spec/configuration_spec.py
    python spec/event_parameters_spec.py
    python spec/event_spec.py
    python spec/guard_spec.py
    python spec/individuation_spec.py
    python spec/one_event_multiple_transitions_spec.py
    python spec/state_action_spec.py
    python spec/state_spec.py
    python spec/transition_action_spec.py
    python spec/boolean_state_getters_spec.py
  )
}

package() {
  cd fluidity
  python -m installer --destdir="$pkgdir" dist/*.whl
  install -Dm644 LICENSE -t "$pkgdir"/usr/share/licenses/$pkgname/
}
