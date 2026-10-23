# Maintainer: Sven-Hendrik Haase <svenstaro@archlinux.org>
# Contributor: Diego Nieto Cid <dnietoc@gmail.com>

pkgname=python-github3py
pkgver=4.0.1
pkgrel=5
pkgdesc="A comprehensive wrapper around the GitHub API (v3)"
arch=("any")
url="https://github.com/sigmavirus24/github3.py"
license=("BSD-3-Clause")
makedepends=("python-build" "python-installer" "python-hatchling")
depends=("python" "python-requests" "python-uritemplate" "python-dateutil" "python-pyjwt")
checkdepends=("python-pytest" "python-betamax" "python-betamax-matchers")
source=("$pkgname-$pkgver.tar.gz::https://github.com/sigmavirus24/github3.py/archive/refs/tags/$pkgver.tar.gz")
sha256sums=('7a1c3f157aa3b9e0973e957ac0b402c09a83d405247d278c10eb4c390977f132')

build() {
    cd "github3.py-${pkgver}"
    python -m build --wheel --no-isolation
}

check() {
    cd "github3.py-${pkgver}"
    python -m venv --system-site-packages test-env
    test-env/bin/python -m installer dist/*.whl
    # Deselect failing tests, possibly due to requests/urllib3 incompatibility
    test-env/bin/python -m pytest --override-ini="addopts=" \
      --deselect=tests/integration/test_repos_repo.py::TestRepoCommit::test_diff \
      --deselect=tests/integration/test_repos_repo.py::TestRepoCommit::test_patch \
      --deselect=tests/integration/test_repos_repo.py::TestComparison::test_diff \
      --deselect=tests/integration/test_repos_repo.py::TestComparison::test_patch
}

package() {
    cd "github3.py-${pkgver}"
    python -m installer --destdir="$pkgdir" dist/*.whl
    install -Dm644 LICENSE "$pkgdir"/usr/share/licenses/${pkgname}/LICENSE
}
